// 一些独立的程序包
package code

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 简单的聊天对话程序
func SimpleChat () {
	// 准备从标准输入读取数据
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input your name: ")

	// 读取数据直到碰到 \n 为止
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("error code: %s\n", err)
	} else {
		// 切片删除最后的\n 切片取出元素不包含结束位置对应的索引
		name := input[:len(input)-1]
		fmt.Printf("Hello, %s! what can i do for you ?\n", name)
	}

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("error code: %s\n", err)
			continue
		}
		// 删除输入字符的回车\n
		input = input[:len(input)-1]

		// 转换为小写
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("bye bye!")
			// 退出当前程序
			os.Exit(0)
		default:
			fmt.Println("Sorry, I ditn't catch you")
		}
	}
}

/* 运行效果
Please input your name:
go
Hello, go! what can i do for you ?
a piece of cake
Sorry, I ditn't catch you
bye
bye bye!
*/