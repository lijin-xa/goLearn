// 函数实现接口
package function

import "fmt"

// 调试器接口
type Invoker interface {
	Call(interface{})
}

// 结构体实现接口
// 结构体类型
type Struct struct {
}

/*
* 实现Invoker的Call方法
* @param s 调用者 （结构体指针类型）
* call 主要实现 打印struct 和 传来的 interface{}类型（void*）的值
*/
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数定位为类型
type FuncCaller func(interface{})

// 实现Invoker的Call
func (f FuncCaller) Call(p interface{}) {

	// 调用f函数本体
	f(p)
}

// 入口函数
func StartInterface() {
	// 声明接口变量
	var invoker Invoker

	// 实例化结构体  new创建的内存空间位于heap, 跟C++不同的是 go存在垃圾回收机制会自动回收
	s := new(Struct)

	// 将实例化的结构体赋值到接口
	invoker = s

	// 使用接口调用实例化结构体的方法
	invoker.Call("hello")

	// 将匿名函数强制转换为FuncCaller类型
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})
	//fmt.Println("is here")

	// 使用接口调用实例化函数的方法 调用函数
	// 相当于变相调用调用者（匿名函数）函数本体
	// f(p) == fmt.Println("from function", v(hello))
	invoker.Call("hello")
}