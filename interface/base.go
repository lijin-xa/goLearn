package _interface

import "fmt"

//  Go语言接口

/**
* 接口是约定的一种合作协议
* 接口是一种类型，也是一种数据结构，不会暴露所含数据的格式，类型和结构
* 接口类型是对其他类型行为的抽象和概括
* Go语言希望通过一个接口精准描述自己的功能
* 通过多个接口的嵌入和组合将简单的接口扩展为负责的接口
*/

/*
接口声明
type 接口类型名 interface {
	 方法名1(参数列表1) 返回值列表1
	 方法名2(参数列表2) 返回值列表2
	 ...
}
*/
// 接口类型名 一般命名时会在单词后面加er
// 当方法名首字母大写 以及 接口类型名首字母也大写 这个方法可以被接口所在包以外的代码访问
// 参数列表和返回值列表可以省略
type writer interface {
	Write([]byte) error
}

// fmt包 - 将一个对象以字符串的形式展示类似于 lua tostring()
type Stringer interface {
	String() string
}

// 接口被实现的条件
// 1 接口的方法与实现接口的类型方法一致
// 2 接口的所有方法均被实现
// 定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

// 定义文件结构 用于实现接口DataWriter
type file struct {

}

// 实现接口的WriteData方法
// f 接收器 - 方法作用的目标
func (f *file) WriteData(data interface{}) error {
	fmt.Println("write data: ", data)
	return nil
}

func Start() {
	// 实例化结构体file 动态申请堆内存空间
	f := new(file)

	// 声明一个DataWriter接口
	var writer DataWriter

	// 将接口赋值给f 初始化这个接口
	writer = f

	// 调用接口的WriterData方法写入数据
	writer.WriteData("go - interface")

	// 测试多类型-单接口
	//test()

	// 测试接口nil之间的判断
	test1()
}

// 接口的实现通过类型 类型与接口之间的关系
// 1 一个类型可以实现多个接口 理解起来有些抽象
// 2 多个类型可以实现相同的接口

// 服务器接口
type Service interface {
	Start()  // 开启服务
	Log()    // 日志输出
}

// 日志器
type Logger struct {
}

// 游戏服务
type GameService struct {
	Logger   // 嵌入日志器
}

// 同一个接收器g
func (g *Logger) Log(l string) {
	//fmt.Println(log)
}

func (g *GameService) Start() {
	fmt.Println("service start")
}

func test() {

	// 实例化GameService 赋值给Service接口

	//TODO
	/* 这里为什么无法还继续提示没有实现 Log方法呢？
	var s Service = new(GameService)
	s.Start()
	s.Log("hello")
	*/
}

/**
* 接口的nil的判断
* 接口的底层实现 由type和data实现
* 如果 data为nil时 type 接口类型仍存在 在判断 == nil 仍然返回false
*/
type MyImplement struct {

}

// 实现fmt.Stringer的String方法
func (m *MyImplement) String() string {
	return "hell - MyImplement"
}

func GetStringer() fmt.Stringer {
	var s *MyImplement = nil

	/*
	想要规避这种情况 发现nil指针时直接返回nil
	if s == nil {
		return nil
	}
	*/

	return s
}

func test1() {
	if GetStringer() == nil {
		fmt.Println("interface nil")
	} else {
		fmt.Println("interface type exist", GetStringer())
	}
}