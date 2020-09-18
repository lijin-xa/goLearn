/*
函数基础
func 函数名(形式参数列表) (返回值类型) {
	函数体
}
*/
package function

import (
	"flag"
	"fmt"
)

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
	//anonymous()

	variableArgs(1, 2, 3)
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

/* 可变参数 即传入的参数可数是可变的
* variableArgs 接受不定数量的参数 参数类型均为int型
* ...type 函数的参数类型存在 - 只能为最后一个参数 == []type
* 两者之间的区别在于
* []type 调用者传递实参时需要传递[]int{} - 数组切片
* ...type 则相对自由，传递任意数量的int型参数即可 variableArgs(1, 2, 3, ...)
*/
func variableArgs(args ...int) {
	for k, v := range args {
		fmt.Println(k, v)
	}

	variableArgsAny(1, "go - varibale", 1.23)
}


// 使用interface{} 传递任意类型的可变参数
func variableArgsAny(args ...interface{}) {
	for _, arg := range args {

		// arg.(type) - 获取到arg类型 只能在switch内配合case使用
		switch arg.(type)  {
		case int:
			fmt.Println(arg, "is int type")
		case string:
			fmt.Println(arg, "is string type")
		default:
			fmt.Println(arg, "is unknown type at now")
		}
	}

}




