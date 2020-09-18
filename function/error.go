/*
* 本节内容
* 错误处理
* 宕机（panic）宕机恢复(recover)
* 计算函数的运行时间
*/
package function

import (
	"fmt"
	"runtime"
	"time"
)

// error 是Go系统声明的定义格式
// Error() 返回错误的具体描述
type error interface {
	Error() string
}

// errors 包
// 自定义一个错误 创建错误对象 （错误字符串一般相对固定，一般在包作用域声明）
func New(text string) error {
	// 返回的错误字符串结构体指针 调用Error() 方法可以直接返回它指向的结构体内部的s字符串的错误信息
	return &errorString{text}
}

// 错误字符串
type errorString struct {
	s string
}

// 返回发生何种错误
func (e *errorString) Error() string {
	return e.s
}

// 宕机 panic （程序运行终止）
func StartPanic() {
	// defer 延迟语句在 panic 引发宕机后执行（先延迟的后执行）
	defer fmt.Println("crash need make thing-2")
	defer fmt.Println("crash need make thing-1")
	// 造成程序崩溃
	panic("program crash")
}

/*
* 宕机恢复 recover()
* 如果当前goroutine 陷入恐慌（宕机），调用recover可以捕获到panic的输入值，并且恢复正常执行
* 类似于try/catch 机制捕获异常
* recover 仅在 defer中有效
*/
type panicContext struct {
	function string
}

func ProtectRun(entry func()) {
	// 这个是什么？ defer func() {}()
	defer func() {
		// 收到panic 传递上下文信息
		err := recover()

		// 错误类型
		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error: ", err)
		default:    // 手动panic错误
			fmt.Println("error: ", err)
		}
	}()

	entry()
}

// 宕机恢复
func StartReCover() {
	// 格式化输出
	// fmt.Sprintf(格式化样式, 格式化参数...)
	// 获取当前时间
	start := time.Now()

	// %#v 输出go语言语法格式的值
	// %+v 对结构体字段名和值进行展开
	fmt.Printf("startTime: %+v\n", start)

	// 程序从start到这里持续的时间 - 计算函数执行时间 返回毫秒（ms）
	// time.Since(start) == time.Now().Sub(start)
	elapsed := time.Since(start)
	fmt.Println("elapsed: ", elapsed)

	// 手动panic
	ProtectRun(func() {
		fmt.Println("手动宕机前")

		// 使用panic传递上下文 cover可以获取到
		//panic(&panicContext{"手动触发 panic"})
		panic("手动触发 panic")

		fmt.Println("手动宕机后")
	})

	// 程序崩溃
	ProtectRun(func() {
		fmt.Println("程序崩溃宕机前")

		var a *int
		*a = 1

		fmt.Println("程序崩溃宕机后")
	})
}

