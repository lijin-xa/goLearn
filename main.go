/*
* 入口函数 main
* 同路径下 文件 - packageName - 必须一致
* package 命名规则
* 1 包名一般小写（简单且有意义）
* 2 包名一般和所在目录名相同，包名不能包含_符号
* 3 一个源文件夹下的所有源码文件只能属于同一个包
* 调用其他包的方法 - packageName.funcName
* 单个文字的方法如果想外部调用 首字母必须大写（相当于publick） 内部使用（private）
* 单个文件下 运行时会自动调用 init()函数  可以有同名的init函数 会以此按顺序调用
*/
package main

import (
	"test/function"
	_interface "test/interface"
	_struct "test/struct"
)

//func init() {
//	fmt.Println("sorry my first come out")
//}

func main() {
	callInterface()
	//callStruct()
	// 输出编写第一个
	//code.PrintGo()
	//num := add.Add(3, 5)
	//fmt.Println(num)
	//fmt.Println(mul.Mul(3, 5))
	//add.Test()
}

// 调用function包
func callFunction() {
	function.Start()
	function.StartInterface()
	function.StartClosure()
	function.StartPanic()
	function.StartReCover()
}

// 调用 _struct包
func callStruct() {
	//_struct.Start()
	_struct.StartOOP()
}

// 调用 interface包
func callInterface() {
	//_interface.Start()
	_interface.StartSort()
}