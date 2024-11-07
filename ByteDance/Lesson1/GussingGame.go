package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 1000
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// fmt.Println("The secret number is", secretNumber)

	fmt.Println("Please input your guess")
	reader := bufio.NewReader(os.Stdin) // 将系统输入转换为 bufio

	for {
		input, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println("error ouccered , try again!")
			continue
		}

		input = strings.Trim(input, "\r\n")

		guess, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Invalid Number. Please input a number!")
			continue
		}

		fmt.Println("You guess is", guess)

		if guess > secretNumber {
			fmt.Println("Guess too big!")
		} else if guess < secretNumber {
			fmt.Println("Guess too small!")
		} else {
			fmt.Print("You are right!")
			break
		}

	}

}
