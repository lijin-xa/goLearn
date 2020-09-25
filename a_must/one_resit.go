package main

import (
	"fmt"
	"math/rand"
)

// 针对第一个问题做回溯
// 1 mongodb 操作
// 2 接口如何使用
// 3 切片的创建
// 4 随机函数
// 5 排序 - QuickSort()

// 接口

// 声明一个接口
type name interface {
	String ()
}

// 用于实现接口
type jam struct {
}

func (j *jam) String () {
	fmt.Println("It is implement interface - String function.")
}

func main() {
	// 实现一个接口
	//interfaceDemo()

	//sliceDemo()

	var number = []int{3, 1, 4, 2, 5}
	quickSort(number)
}

func interfaceDemo() {
	// 声明接口并赋值为 *jam 类型
	// 虽然左右值类型不同，但由于j是一个接口类型，并且jam实现器已经实现了接口内的String方法 所以可以赋值
	//var j name = new(jam)
	var j name = &jam{}
	j.String()
}

// 切片的熟练掌握
// 声明一个100长度的切片
func sliceDemo() {
	//var numList = []int{}
	// 动态创建一个切片
	var numList = make([]int, 100)

	// 插入100个元素
	for i := 0; i < 100; i++ {
		numList[i] = rand.Intn(100)
	}

	fmt.Println(len(numList))
}

// 快速排序
func quickSort(numList []int) {

	// 步骤1
	// 将 3 1 4 2 5  固定3  1 2 3 4 5
	// 2 3 1 4 5

	// 比较位
	var compareV = numList[0]
	for i := len(numList)-1; i > 0; i-- {
		if numList[i] < compareV {
			//lower := numList[i]
			// 先将这个数 从切片中删除
			//numList = append(numList[:i], numList[i+1:]...)

			// 再插入到第一个位置
		}
	}
}

//type error interface {
//	Error() string
//}
//
//// 需要一个类型实现接口
//type RPCError struct {
//}
//
//func (r *RPCError) Error() string {
//	fmt.Println("this is my interface")
//}
//
//func Start() {
//	var rpcErr = new(RPCError)
//	rpcErr.Error()
//}

