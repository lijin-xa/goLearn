package _struct

import (
	"fmt"
	"math"
)

/**
* 通过struct 实现面向对象编程
* Go 没有类(class) 的概念，也不支持类的继承等面向对象的概念
* 但是通过结构体(struct)来实现面向对象编程
* 结构体内嵌实现类的派生
* 接收器(receiver)实现 - 结构体添加方法
*/

// 入口接口
func StartOOP() {
	baseTypeAddMethod()

	variableInAssigment()

	innerStruct_Type()

	OOPTest()

	// 结构体内嵌初始化测试
	InitStruct()

}

type Vec2 struct {
	X, Y float32
}

// 加
func (v Vec2) Add(other Vec2) Vec2 {
	return Vec2{
		v.X + other.X,
		v.Y + other.Y,
	}
}

// 减
func (v Vec2) Sub(other Vec2) Vec2 {
	return Vec2{
		v.X - other.X,
		v.Y - other.Y,
	}
}

// 乘
func (v Vec2) Scale(s float32) Vec2 {
	return Vec2{v.X * s, v.Y * s}
}

// 距离
func (v Vec2) DistanceTo(other Vec2) float32 {
	dx := v.X - other.X
	dy := v.Y - other.Y

	return float32(math.Sqrt(float64(dx * dx + dy * dy)))
}

// 单位化
func (v Vec2) Normalize() Vec2 {
	mag := v.X * v.X + v.Y * v.Y
	if mag > 0 {
		oneOverMag := 1 / math.Sqrt(float64(mag))

		// Invalid operation: v.Y * oneOverMag (mismatched types float32 and float64)
		// float32(x) 无法 float64(y) 无法一起操作？
		return Vec2{v.X * float32(oneOverMag), v.Y * float32(oneOverMag)}
	}

	return Vec2{0, 0}
}

// TODO - 玩家矢量移动

// 基本类型添加方法
type MyInt int

func (m MyInt) IsZero() bool {
	return m == 0
}

func (m MyInt) Add(other int) int {
	return other + int(m)
}

// 测试基础类型添加方法
func baseTypeAddMethod() {
	var b MyInt  // 默认为0
	fmt.Println(b.IsZero())

	b = 1
	fmt.Println(b.Add(1))
}

// 无论是普通函数还是结构体方法，只要他们的签名一致
// 就可以使用与它们签名一致的函数变量来保存普通函数和结构体的方法
type class struct {

}

func (c *class) Do(v int) {
	fmt.Println("call struct method: ", v)
}

func funcDo(v int) {
	fmt.Println("call function:", v)
}

// 测试这两种情况的下给签名一致的变量赋值
// 结果显示 只要函数变量与其签名一致 即可调用
func variableInAssigment() {
	// 声明一个函数回调
	var delegate func(int)

	// 创建一个结构体实例
	c := new(class)

	// 将回调设为c的Do方法
	delegate = c.Do

	delegate(100)

	// 将回调设为funcDo函数
	delegate = funcDo

	delegate(100)
}

// TODO 事件系统
// 1 能够实现事件的一方，可以根据事件ID或者名字注册对应的事件
// 2 事件发起者，会根据注册信息通知这些注册者
// 3 一个事件可以有多个实现方响应


/*
* 类型内嵌和结构体内嵌 - 模拟类的继承
* 结构体内的匿名字段（结构体类型）- 结构体可以包含内嵌结构体
* 结构体的内嵌特性:
* 1 内嵌结构体可以直接访问其成员变量 - 无须一层层方法 可以直接在外层访问内层的字段
* 需要注意的是 如果结构体内部存在重名字段 就必须一层一层访问赋值 避免歧义
* 2 内嵌的结构体字段名就是它的类型名
*/

type innerS struct {
	int1 int
	int2 int
}

type outerS struct {
	 b int
	 c float32
	 int     // anonymous field
	 innerS  // anonymous field
}

func innerStruct_Type() {

	// 创建一个outerS的实例化
	outer := new(outerS)
	outer.b = 1
	outer.c = 1.1

	// outer.int可以获取到匿名字段赋值 同样也可以获取到对应的值
	outer.int = 2
	outer.int1 = 3
	outer.int2 = 4
	fmt.Printf("%+v\n", outer)
}

// 结构体内嵌 - 模拟类的继承 简单的实例
// 通过结构体内嵌特性实现多重继承
type Flying struct {
}

func (f *Flying) Fly(desc string) {
	fmt.Println(desc, " can fly")
}

type Walkable struct {
}

func (w *Walkable) Walk(desc string) {
	fmt.Println(desc, " can walk")
}

// 人类
type Human struct {
	desc string
	Walkable
}

// 鸟类
type Bird struct {
	desc string
	Flying
	Walkable
}

func OOPTest() {
	b := new(Bird)
	b.desc = "bird"

	b.Fly(b.desc)
	b.Walk(b.desc)

	h := new(Human)
	h.desc = "human"
	h.Walk(h.desc)
}

// 初始化内嵌结构体
type Wheel struct {
	Size int
}

type Car struct {
	Wheel

	// 内嵌结构化的声明在内部 在初始化时必须带入其声明字段
	Engine struct {
		Power int   // 功率
		Type string // 类型
	}
}

func InitStruct () {
	c := Car {
		Wheel: Wheel {
			Size: 1,
		},

		// 在初始化时还是必须带入声明时的定义
		Engine: struct {
			Power int
			Type string
		}{
			Type: "1.5T",  // 分隔符,必须加上
			Power: 143,
		},
	}

	fmt.Printf("car %+v\n", c)
}