package _package
// os包 提供了操作系统函数的接口

import (
	"fmt"
	"os"
	"os/exec"
	user2 "os/user"
)

// 入口函数
func StartOs() {
	name, _ := os.Hostname()
	fmt.Println(name)

	// os/exec包练习
	execTest()

	// os/user包练习
	userTest()
}

// os/exec 执行外部命令
func execTest() {
	// 在环境变量指定的目录下搜索可执行文件
	f, err := exec.LookPath("main")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(f)
}

// os/user 获取当前用户信息
func userTest() {
	// user 是一个结构体
	user, _ := user2.Current()
	fmt.Println(user.Username)
	fmt.Println(user.HomeDir)       // 用户主目录
}

// os/signal 信号处理
