package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义模型
type User struct {
	ID       uint   `gorm:"primaryKey"` // 主键
	Name     string `gorm:"size:100"`   // 用户名字段
	Email    string `gorm:"unique"`     // 邮箱字段，唯一约束
	Password string `gorm:"size:255"`   // 密码字段
}

func main() {
	// 数据库连接配置
	dsn := "root:Zq123456@tcp(127.0.0.1:3306)/GormTest?charset=utf8mb4&parseTime=True&loc=Local"

	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接数据库失败: %v", err)
	}

	fmt.Println("成功连接到 MySQL 数据库！")

	// 自动迁移，创建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("自动迁移失败: %v", err)
	}
	fmt.Println("表创建成功！")

	// 插入数据
	newUser := User{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password123",
	}

	if err := db.Create(&newUser).Error; err != nil {
		log.Fatalf("插入数据失败: %v", err)
	}
	fmt.Printf("成功插入用户: %+v\n", newUser)

	// 查询数据
	var user User
	if err := db.First(&user, "email = ?", "alice@example.com").Error; err != nil {
		log.Fatalf("查询数据失败: %v", err)
	}
	fmt.Printf("查询到用户: %+v\n", user)
}
