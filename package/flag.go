package _package
// 命令行参数解析

import (
	"flag"
	"fmt"
	"strings"
)

// flag.Type() *Type 返回一个相应类型的指针
// go run main.go -name "name" -age=xx 获取命令行数据 没有输入相应参数-就使用设置的默认值
var Input_name = flag.String("name", "gerry", "input your name")
var Input_age = flag.Int("age", 20, "input your age")
var Input_flagvar int

func Init() {
	// 将flag绑定到一个变量上
	flag.IntVar(&Input_flagvar, "flagname", 1024, "help message for flagname")
}

func StartFlag() {
	//Init()
	//// 解析命令行参数
	//flag.Parse()
	//
	//fmt.Println("name = ", *Input_name)
	//fmt.Println("age = ", *Input_age)
	//fmt.Println("flagname = ", Input_flagvar)

	defineTest()
}

// 实现flag.Value接口实现自定义flag  receiver接收器是指针类型
type sliceValue []string

type Value interface {
	String() string
	Set(string) error
}

func newSliceValue(vals []string, p *[]string) *sliceValue {
	*p = vals
	return (*sliceValue)(p)
}

func (s *sliceValue) Set(val string) error {
	*s = sliceValue(strings.Split(val, ","))
	return nil
}

func (s *sliceValue) String() string {
	*s = sliceValue(strings.Split("default is me", ","))
	return "It's none of my business"
}

// TODO 还是有一些需要加深理解的地方
func defineTest() {
	var languages []string

	// new一个存放命令行参数值的slice - languages转换为sliceValye类型
	flag.Var(newSliceValue([]string{}, &languages), "slice", "I like programming `languages`")
	flag.Parse()

	//打印结果slice接收到的值
	fmt.Println(languages)
}