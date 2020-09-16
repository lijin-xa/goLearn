/*
* 基础语法规则 条件语句等等
*/
package main

import (
	"fmt"
	"math"
)

// Go语言而言 左括号必须和关键字在一行
func main() {

	if true {
		fmt.Println("this is my first condition syntax")
	// 前段代码块的右括号必须和关键字 else/else if 在同一行
	} else {
		fmt.Println("no some")
	}

	// if条件编译 可以增加一个执行语句
	if err := Connect(); err != nil {
		fmt.Println("have something", err)
		return
	} else {
		fmt.Println("have nothing", err)
	}

	// 基础遍历
	LoopSyntax()

	// 嵌套遍历
	MultiTable()

	// key-value 遍历
	keyValueIterator()

	//switch case
	switchSyntax()

	//额外练习
	//mathMethod()
}

/**
* go uint类型数据的操作问题
*/
func mathMethod() {
	const num = 5

	var count uint32         //uint32 默认0
	fmt.Println("count: ", count)                //0

	// 输出未知 - 非-1
	fmt.Println("count - 1: ", count - 1)
	fmt.Println("float64(count - 1): ", float64(count - 1))
	consume := num * math.Pow(2, float64(count - 1))

	// 会输出uint32的默认值 0
	fmt.Println("uint32(consume): ", uint32(consume))

	// 输出未知
	fmt.Println("consume: ", num * math.Pow(2, float64(count - 1)))
}

// interface{} == void*
func Connect() interface{} {
	return nil
}

// for 一些特殊格式 冒泡排序
func LoopSyntax() {

	// 不支持while do-while语句
	sum := 1
	// for i := 1; i < 10; i++ {
	// }

	i := 1
	for ; ; i++ {
		sum = sum + 1
		fmt.Println("sum: ", sum)
		if sum >= 10 {
			break
		}
	}

	/*
	相当于变相的do - while
	for {
		if condition {
			break
		}
	}
	*/

	// var arr []int = []int {1,3,5,2,4}
	// 简单的冒泡排序
	arr := []int {1,3,5,2,4}
	for i := 1; i != 5; i++ {
		for j := 1; j != 5 - i; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)
}

// 9*9乘法表
func MultiTable() {

	outerLoop:
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf(" %d * %d = %d", i, j, i*j)

			if i == 8 && j == 7 {

				// 加入标签后 可以直接跳出标签外层 - 程序结束
				break outerLoop

				// 跳出这一次循环 开始下一次循环
				//continue outerLoop
			}
		}

		//换行
		fmt.Println()
	}
}

// range 实现 key - value遍历
func keyValueIterator() {

	str := []string {"go", "syntax", "iterator"}
	for k, v :=  range str {
		fmt.Println(k, v)
	}

	var str1 = "hello world"
	for k, v := range str1 {
		fmt.Println(k, v)          // 输出是ASCII码
	}

	// map 成员 key : value
	map1 := map[string]int {
		"go" : 1,
		"c++" : 2,
		"lua" : 3,
	}
	for k, v := range map1 {
		fmt.Println(k, v)
	}
}

/*
* switch case 语句
* case 是一个独立的代码块 并不会像C语言那样执行完毕紧接着执行下一个case
* fallthrough 可以兼容一些移植代码 - 具体使用未有概念
*/
func switchSyntax() {

	var temp = "go"
	switch temp {

	// 存在一个符合的条件均可进入case代码块
	case "go", "c++" :
		fmt.Println("go + c++")
	default:
		fmt.Println("go")
	}
}


