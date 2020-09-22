package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

// go并发 - 多线程编程

// go 语言的并发通过goroutine特性完成 goroutine类似于线程
// goroutine 是有Go语言运行时调度完成 线程是由操作系统调度完成
// channle 在多个goroutine间通信

// 并发/并行
// 多线程程序在单核心的CPU上运行，称为并发 多线程程序在多核心的CPU上运行，称为并行
// goroutine并行 也即 一个goroutine占用一个Cpu逻辑处理器
// 并发 同事管理很多事情，这些事情可能只做了一半就被暂停去做别的事情 - cpu的时间片被占用了

// go关键字创建goroutine
// 将go 声明放到一个需要调用的函数之前，在相同地址空间调用运行这个函数
// 这样函数执行时便会作为一个独立的并发线程 - 这个线程在go的语法中叫做goroutine

// channle Go语言在语言级提供的goroutine间的通讯方式
// channle 传递对象的过程和函数传递参数比较一致

// goroutine Go语言中的轻量级线程实现，由Go运行时（runtime）管理，Go程序会智能地将goroutine中的任务分配给每一个CPU
// go 函数名(参数列表)  返回值会被忽略
// 使用channel把数据从goroutine中返回出来

func Start() {

	// goroutine - 归属于running函数
	go running()

	var input string

	// 输入数据 goroutine - Start()函数
	fmt.Scanln(&input)
	fmt.Println("input data: ", input)

	// Cpu核心数量
	num := runtime.NumCPU()
	fmt.Println("Cpu core num: ", num)

	// runtime.GOMAXPROCS(逻辑CPU数量)
	// <1 - 不修改任何数值
	// =1 - 单核心执行
	// >1 - 多核并发执行
	// 从Go 1.5开始 默认执行 - 多核并发执行
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)

		// 延时一秒
		time.Sleep(time.Second)
		if times > 2 {
			break
		}
	}
}

/*
使用匿名函数创建goroutine
go func() {
	函数体
}(调用参数列表)
调用参数列表 - 启动goroutine时需要向匿名函数传递的调用参数 （不可省略）
*/

// go 使用消息队列作(channle)为通讯方式 而非使用共享内存
// 消息机制认为每个并发单元是自包含的，独立的个体，并且都有自己的变量，但在不同并发单元间这些变量不共享

// go build -race命令 检测资源竞争的功能
// 资源竞争 - 多线程常用的办法 - 原子操作（atomic） 加锁 mutex

// 原子操作 线程在执行期间 不会发生上下文切换
// 互斥锁 - 在代码上创建一个临界区，保证同一时间只有一个goroutine可以执行临界区代码

// channle 通道 - 任何时候只能有一个goroutine访问通道进行发送数据和获取数据
func channleTest() {
	// var 通道变量 chan 通道类型
	// 声明一个int型通道ch1
	var ch1 chan int
	// 通道是引用类型，必须使用make进行创建
	ch1 = make(chan int)
	// 单向通道 - 只读 <-chan 只写 chan<-
	// var onlyRead <-chan int = ch1
	// var onlyWrite chan<- int = ch1

	type data struct {
	}

	// 创建data指针类型的通道
	// ch2 := make(chan *data)

	// <- 操作符 向通道发送数据 发送数据后将阻塞直到数据被接受
	ch1 <- 1

	// 使用通道接受数据 每次接收只能接收一个数据元素
	// 阻塞模式接收数据 - 直到发送方发送数据
	num := <-ch1
	// 非阻塞接收数据 ok 表示是否接收数据
	// num, ok = <- ch1
	// 接收任意数据，忽略接收的数据
	<-ch1
	// 循环接收
	for num = range ch1 {
	}

	fmt.Println(num)

	// 关闭通道  使用非阻塞接收方式可以判断管道是否关闭
	close(ch1)
}
