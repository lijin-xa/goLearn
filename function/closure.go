package function

import "fmt"

/*
* 闭包 - 能够读取其他函数内部变量的函数
* 可以理解为定义在函数内部的函数
* 闭包具备一定封装行， 无法在函数外部直接访问或修改闭包内部的变量
 */
func Closure() {
	str := "hello world"

	foo := func() {

		str = "hello dube"
		fmt.Println(str)
	}

	fmt.Println(str)
	foo()
}

// 返回的是一个闭包函数 func() int {}
func Accumulate(value int) func() int {

	return func() int {
		value++
		return value
	}
}

// 入口函数
func StartClosure() {

	// accumulator func()int 的函数变量
	accumulator := Accumulate(1)

	//
	fmt.Println(accumulator())

	fmt.Println(accumulator())

	fmt.Printf("%p\n", &accumulator)

	accumulator2 := Accumulate(10)

	fmt.Println(accumulator2())
	fmt.Printf("%p\n", &accumulator2)
}

