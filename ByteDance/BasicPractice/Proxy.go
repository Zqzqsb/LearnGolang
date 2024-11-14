package main

import (
	"bufio"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

// socks5 的协议控制相关
const (
	socks5Ver  = 0x05
	cmdConnect = 0x01
	atypeIPV4  = 0x01
	atypeHOST  = 0x03
	atypeIPV6  = 0x04
)

func main() {
	// 启动混合代理，监听 127.0.0.1:1080
	server, err := net.Listen("tcp", "127.0.0.1:1080")
	if err != nil {
		log.Fatalf("Failed to start mixed proxy: %v", err)
	}
	log.Printf("Mixed proxy listening on %s", "127.0.0.1:1080")

	// 监听事件的循环
	for {
		client, err := server.Accept()
		if err != nil {
			log.Printf("Accept failed: %v", err)
			continue
		}
		go processMixedProxy(client)
	}
}

func processMixedProxy(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	firstByte, err := reader.Peek(1)
	if err != nil {
		log.Printf("Fail to read first byte: %v ", err)
		return
	}

	// 读取第一个字节，判断协议类型
	switch firstByte[0] {
	case socks5Ver:
		log.Printf("Handling socks5 request from %v", conn.RemoteAddr())
		if err := processSocks5(reader, conn); err != nil {
			log.Printf("Socks5 handling failed: %v", err)
		}
	case 'C':
		log.Printf("Hadling http request from %v", conn.RemoteAddr())
		if err := processHTTP(reader, conn); err != nil {
			log.Printf("HTTP handling failed: %v", err)
		}
	default:
		log.Printf("Unsupported protocol from %v", conn.RemoteAddr())
	}
}

func processSocks5(reader *bufio.Reader, conn net.Conn) error {
	if err := auth(reader, conn); err != nil {
		return fmt.Errorf("Socks5 auth failed: %w", err) // 返回一个新包装的err
	}
	if err := connect(reader, conn); err != nil {
		return fmt.Errorf("Socks5 connect failed: %w", err)
	}
	return nil
}

// Socks5的认证逻辑
func auth(reader *bufio.Reader, conn net.Conn) error {
	// +----+----------+----------+
	// |VER | NMETHODS | METHODS  |
	// +----+----------+----------+
	// | 1  |    1     | 1 to 255 |
	// +----+----------+----------+
	// VER: 协议版本，socks5为0x05
	// NMETHODS: 支持认证的方法数量
	// METHODS: 对应NMETHODS，NMETHODS的值为多少，METHODS就有多少个字节。RFC预定义了一些值的含义，内容如下:
	// X’00’ NO AUTHENTICATION REQUIRED
	// X’02’ USERNAME/PASSWORD

	_, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read ver failed: %w", err)
	}

	methodSize, err := reader.ReadByte()
	if err != nil {
		return fmt.Errorf("read methodSize failed: %w", err)
	}

	method := make([]byte, methodSize)
	_, err = io.ReadFull(reader, method)
	if err != nil {
		return fmt.Errorf("read method failed: %w", err)
	}

	// +----+--------+
	// |VER | METHOD |
	// +----+--------+
	// | 1  |   1    |
	// 认证方法相关
	// 告知客户端无须认证
	_, err = conn.Write([]byte{socks5Ver, 0x00})
	if err != nil {
		return fmt.Errorf("write failed: %w", err)
	}

	return nil
}

// socks5
func connect(reader *bufio.Reader, conn net.Conn) error {
	// 解析请求头
	// +----+-----+-------+------+----------+----------+
	// |VER | CMD |  RSV  | ATYP | DST.ADDR | DST.PORT |
	// +----+-----+-------+------+----------+----------+
	// | 1  |  1  | X'00' |  1   | Variable |    2     |
	// +----+-----+-------+------+----------+----------+
	// VER 版本号，socks5的值为0x05
	// CMD 0x01表示CONNECT请求
	// RSV 保留字段，值为0x00
	// ATYP 目标地址类型，DST.ADDR的数据对应这个字段的类型。
	//   0x01表示IPv4地址，DST.ADDR为4个字节
	//   0x03表示域名，DST.ADDR是一个可变长度的域名
	// DST.ADDR 一个可变长度的值
	// DST.PORT 目标端口，固定2个字节

	buf := make([]byte, 4)
	_, err := io.ReadFull(reader, buf)
	if err != nil {
		return fmt.Errorf("read header failed: %w", err)
	}

	ver, cmd, atype := buf[0], buf[1], buf[3]
	if ver != socks5Ver || cmd != cmdConnect {
		return fmt.Errorf("unsupported ver or cmd")
	}

	var addr string

	/*
		ATYP (buf[3])：地址类型，决定目标地址的格式。可能的值有：
		•	0x01：IPv4 地址（4 字节）
		•	0x03：域名地址（第一个字节为长度，后面为域名）
		•	0x04：IPv6 地址（16 字节）
		读出 ip 地址 或者 域名
	*/
	switch atype {
	case atypeIPV4:
		_, err = io.ReadFull(reader, buf)
		if err != nil {
			return fmt.Errorf("read ipv4 address faile: %w", err)
		}
	case atypeHOST:
		hostSize, err := reader.ReadByte()
		if err != nil {
			return fmt.Errorf("read host size failed: %w", err)
		}
		host := make([]byte, hostSize)
		_, err = io.ReadFull(reader, host)
		if err != nil {
			return fmt.Errorf("read host failed: %w", err)
		}
		addr = string(host)
	case atypeIPV6:
		return errors.New("IPv6 not supported yet")
	default:
		return errors.New("unsupported address type")
	}

	// 读取最后两个字段的 port
	_, err = io.ReadFull(reader, buf[:2])
	if err != nil {
		return fmt.Errorf("read port failed: %w", err)
	}
	port := binary.BigEndian.Uint16(buf[:2])

	dest, err := net.Dial("tcp", fmt.Sprintf("%v:%v", addr, port))

	if err != nil {
		return fmt.Errorf("dial dst failed:%w", err)
	}
	defer dest.Close()

	// 可以访问目标地址 返回成功
	_, err = conn.Write([]byte{0x05, 0x00, 0x00, 0x01, 0, 0, 0, 0, 0, 0})
	if err != nil {
		return fmt.Errorf("write response failed: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		_, _ = io.Copy(dest, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, dest)
		cancel()
	}()

	<-ctx.Done()
	return nil
}

// 处理 HTTPS 代理连接
func processHTTP(reader *bufio.Reader, conn net.Conn) error {
	// 读取客户端的 CONNECT 请求
	request, err := reader.ReadString('\n')
	if err != nil {
		return fmt.Errorf("failed to read HTTPS request: %w", err)
	}

	// 解析 CONNECT 请求
	requestParts := strings.Fields(request)
	if len(requestParts) < 3 || strings.ToUpper(requestParts[0]) != "CONNECT" {
		return fmt.Errorf("invalid HTTPS request: %s", request)
	}
	dest := requestParts[1]
	log.Printf("HTTPS CONNECT to %s", dest)

	// 响应 200 Connection Established
	_, err = conn.Write([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	if err != nil {
		return fmt.Errorf("failed to write HTTPS response: %w", err)
	}

	// 连接到目标服务器
	destConn, err := net.Dial("tcp", dest)
	if err != nil {
		return fmt.Errorf("failed to connect to destination %s: %w", dest, err)
	}
	defer destConn.Close()

	// 开始转发客户端与目标服务器之间的数据
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		_, _ = io.Copy(destConn, reader)
		cancel()
	}()
	go func() {
		_, _ = io.Copy(conn, destConn)
		cancel()
	}()

	<-ctx.Done()
	return nil
}
