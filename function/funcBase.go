package function

import (
	"flag"
	"fmt"
)

/*
func 函数名(形式参数列表) (返回值类型) {
	函数体
}
*/

// 如果一组形参或者返回值类型相同，不必写出没个参数类型
// func add1(x, j int)
// func add2(x int, y int)

// 支持对返回值进行命名 命名的返回值类型为类型的默认值
func Add() (a, b int) {

	a = 1
	b = 2

	// 带命名的返回值 在return时可以省略返回值具体参数 （return a, b == return）
	return
}

// 入口程序
func Start() {
	// 简单的练习
	// test()

	// 匿名函数
	anonymous()
}

// 函数参数传递
func test() {
	type InnerData struct {
		a int
	}

	type Data struct {
		complex []int
		instance InnerData  // 结构体变量
		ptr *InnerData      // 结构体指针
	}

	in := Data {
		complex: []int {1, 2, 3},
		instance: InnerData {
			5,
		},
		ptr: &InnerData{1},
	}

	// 打印结构体 %+v %v 之间的区别
	// %+v 可以打印出结构体内部 - 详细结构 包含字段
	// %v 只会打印出内部的数据
	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in value: %v\n", in)

	// var f func()
	// 函数也是一种类型 可以将函数名作为值复制给变量
	f := test
	f()
}

// 匿名函数（一些简单应用场景）
func anonymous() {
	// flag 包的API简单了解
	// 从命令行输入 --skill 可以将 = 后的字符串传入给skillParam 指针变量
	var skillParam = flag.String("skill", "","skill to perform")

	// 解析命令行参数 skillParam指针变量将指向命令行传入的值
	flag.Parse()

	// map内的value为匿名函数
	var skill = map[string]func() {
		"fire": func() {
			fmt.Println("chicken fire")
		},
		"run": func() {
			fmt.Println("soldier run")
		},
		"fly": func() {
			fmt.Println("angel fly")
		},
	}

	// 从命令行获取到的字符串 定位到map中对应的键值
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}

	// 命令行输入 go run main.go --skill=xx
}




