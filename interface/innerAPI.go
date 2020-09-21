package _interface

import (
	"fmt"
	//"io"
	"sort"
)

// 内置包 - 提供的接口

/*
* 排序 - sort.Interface接口 对任何序列排序的功能
* 序列的长度， 表示两个元素比较的结果， 一种交换两个元素的方式
* 序列的表示经常是一个切片
*/

// 这个接口下的三个方法必须实现 才可以使用
type Interface interface {
	Len() int              // 获取元素数量
	Less(i, j int) bool    // 1，j是序列元素的指数 - 比较元素
 	Swap(i, j int)         // 交换元素
}

// 为了能够让sort识别需要排序的字符串序列 使用排序的字符串序列实现sort.interface接口
type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i, j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func StartSort() {
	var a = 1
	var b = 2

	// go 和 lua一样的，也是先确定右端值，可以这样直接交换数据
	a, b = b, a
	fmt.Println(a, b)

	names := MyStringList{
		"engine",
		"double",
		"cat",
		"banana",
		"apple",
	}

	sort.Sort(names)

	for _, v := range names {
		fmt.Println(v)
	}

	// 便捷排序 单类型排序
	instantSort()

	// 结构体排序
	structSort()

	// 内置接口
	innerInterface()

	// 接口断言转换
	convertInterface()

	// 空接口测试
	emptyInterface()

	// 类型断言 type - switch语句
	interfaceAssertion()

	// 错误接口
	interfaceError()
}

// 常见类型的快捷排序
// 字符串切片排序
func instantSort() {
	// sort - StringSlice类型 已经实现sort.Interface接口方法
	// type StringSlice []string - 按照字符 ASCII 值升序
	names := sort.StringSlice{
		"engine",
		"double",
		"cat",
		"banana",
		"apple",
	}

	// 默认升序排序
	sort.Sort(names)
	fmt.Println(names)

	// 整形切片排序 - type IntSlice []int
	numbers := sort.IntSlice{4,3,3,2,1}
	sort.Sort(numbers)
	fmt.Println(numbers)

	// 直接对字符串切片进行排序
	sort.Strings(names)

	// 浮点数切片排序 - type Float64Slice []flat64
}

type HeroKind int

const (
	None     HeroKind = iota // 枚举自增从0开始
	Tank                     // 坦克1
	Assassin                 // 刺客2
	Mage                     // 法师3
)

// 英雄名单结构体
type Hero struct {
	Name string
	Kind HeroKind
}

// &Hero struct类型的切片  也可以使用 []Hero类型
// 使用指针传递效率更高
type Heros []*Hero

func (s Heros) Len() int {
	return len(s)
}

// 实现sort.Interface的比较元素方法
func (s Heros) Less(i, j int) bool {
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}

	return s[i].Name < s[j].Name
}

// 实现sort.Interface的交换元素方法
func (s Heros) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 结构体数据进行排序
func structSort() {
	heros := Heros{
		&Hero{"吕布", Tank},
		&Hero{"李白", Assassin},
		&Hero{"妲己", Mage},
		&Hero{"貂蝉", Assassin},
		&Hero{"关羽", Tank},
		&Hero{"诸葛亮", Mage},
	}

	// 重新实现Slice方法 - 就可以不用去实现 Len - Less - Swap这三个方法
	// 直接给切片进行排序
	sort.Slice(heros, func(i, j int) bool {
		if heros[i].Kind != heros[j].Kind {
			return heros[i].Kind < heros[j].Kind
		}

		return heros[i].Name < heros[j].Name
	})

	sort.Sort(heros)
	for _, v := range heros {
		fmt.Println(v.Name)
	}
}

/*
* 接口的嵌套
* 一个接口中可以包含一个或多个接口，相当于将内嵌接口的方法列举到外层接口中一样
* 只要接口的方法被实现，就可以被调用
*/

// 声明一个接口的实现器
type device struct {
}

func (d *device) Write(p []byte) (n int, err error) {
	return 0, nil
}

func (d *device) Close() error {
	return nil
}

func innerInterface() {
	// 创建一个写入关闭器的实例
	//var wc io.WriteCloser = new(device)

	// wc可以调用其内嵌接口的方法
	//wc.Write(nil)
	//
	//wc.Close()
}

/**
* 接口与类型的转换
* 使用接口断言 - 将接口转换成另一个接口
* t := i.(T)
* i 接口变量 T 转换的目标类型 t 转换后的变量
* T 接口类型 如果i,T的动态类型满足 则会将i转换为T类型
*/

type Flyer interface {
	Fly()
}

type Walker interface {
	Walk()
}

type bird struct {
}

func (b *bird) Fly() {
	fmt.Println("bird: fly")
}

func (b *bird) Walk() {
	fmt.Println("bird walk")
}

type pig struct {
}

func (p *pig) Walk() {
	fmt.Println("pig walk")
}

func convertInterface() {
	animals := map[string]interface{} {
		"bird": new(bird),
		"pig": new(pig),
	}

	//var obj interface{} = new(bird)

	for _, v := range animals {

		f, isFlyer := v.(Flyer)

		w, isWalker := v.(Walker)

		fmt.Printf("f type: %T = %+v\n", f, f)
		//fmt.Println(f, isWalker)

		if isFlyer {
			f.Fly()
		}

		if isWalker {
			w.Walk()
		}
	}

	// 将接口转换为指针类型
	p1 := new(pig)
	var a Walker = p1
	p2 := a.(*pig)

	fmt.Printf("p1: %T p2: %T\n", p1, p2)
}

// 空接口 - 内部实现保存了对象的类型和指针 interface {}
// 使用空接口保存一个数据的过程会比直接用数据对应类型的变量保存稍慢
// 空接口比较时 不能比较空接口中的动态值
func emptyInterface() {

	var any interface{} = 1

	// 空接口类型不能直接赋值 需要使用类型断言
	// 错误var b int = any 也可以这样 b1 := any
	b := any.(int)

	fmt.Printf("%T\n", any)
	fmt.Println(any)
	fmt.Println(b)

	var c interface{} = []int {1, 2, 3}
	var d interface{} = []int {1, 2, 3}
	fmt.Println(c, d)
	// comparing uncomparable type int[] - 比较无法比较类型
	// fmt.Println(c == d)
	// map, slice 是无法比较的，会panic
	// channel array struct func 是可以比较的
}

// 类型断言 type - switch 语句
type Alipay struct {
}

// Alipay添加方法
func (a *Alipay) CanUseFaceId() {
	fmt.Println("support Use Face")
}

type Cash struct {
}

func (c *Cash) Stolen() {
	fmt.Println("maybe stolen")
}

type ContainCanUseFaceId interface {
	CanUseFaceId()
}

type ContainStolen interface {
	Stolen()
}

func interfaceAssertion() {

	var payMethod interface{} = new(Alipay)
	//var payMethod interface{} = new(Cash)

	// payMethod 的值和类型
	switch payMethod.(type) {
	// 类型是 ContainCanUseFaceId
	case ContainCanUseFaceId:
		fmt.Printf("Alipay - payMethod: %T, %+v\n", payMethod, payMethod)
	case ContainStolen:
		fmt.Printf("Cash - payMethod: %T\n", payMethod)
	}
}

// error接口
type error interface {
	Error() string
}

// 重新定义一个实现器 - 实现接口的Error()
type newError struct {
	Num     float64
	problem string
}

func (n *newError) Error() string {
	// strconv.FormatFloat 将浮点数转化为字符串
	//err := strconv.FormatFloat(n.Num, 'f', 10, 64) + "this is new error"
	return n.problem
}

func testError() (float64, error){
	return 1.0, &newError {Num: 1.0, problem: "this is new error"}
}

// 自定义错误类型
func interfaceError() {
	result, err := testError()

	fmt.Println(result, err)
}

































