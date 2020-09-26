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

	//numList := sliceDemo(10)

	var numList = []int{1, 7, 7, 9, 1, 8, 5, 0, 6, 0}
	lower := 0
	high := len(numList) - 1
	var a = 0
	count := &a

	quickSort(numList, lower, high, count)
	fmt.Println("排序后: ",numList)
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
func sliceDemo(length int) []int {
	//var numList = []int{}
	// 动态创建一个切片
	var numList = make([]int, length)

	// 插入100个元素
	for i := 0; i < length; i++ {
		numList[i] = rand.Intn(length)
	}

	fmt.Println(len(numList))

	return numList
}

// 快速排序
func quickSort(numList []int, lower int, high int, count *int) {

	comparePos := test1(numList, lower, high, count)

	fmt.Println("the first comparePos: ", "lower: ", lower, " high: ", high,
		" comparePos: ",  comparePos)

	// 左端排序
	if lower < comparePos {
		//fmt.Println("lower: ", lower, " high: ", comparePos-1)
		fmt.Println("左端数: ", numList[lower:comparePos-1])
		quickSort(numList, lower, comparePos-1, count)
	}
	// 右端排序
	if high > comparePos {
		//fmt.Println("high: ", high, " lower: ", comparePos+1)
		fmt.Println("右端数: ", numList[comparePos+1:high])
		quickSort(numList, comparePos+1, high, count)
	}
}

func test2(numList []int, lower int, high int, count *int) int {
	*count = *count + 1
	fmt.Println("----------- ", *count," -----------")

	// 比较位
	var compareV = numList[lower]
	var comparePos = lower

	fmt.Println("start", numList,
		"comparePos: ", comparePos, " compareV: ", compareV)

	for i := high; i > lower; i-- {
		// 小于放左端
		if i > comparePos && numList[i] <= compareV {

			fmt.Println("compareV[",comparePos, "]:", compareV,
				"[", i , "]: ", numList[i],)

			// 和比较位交换
			numList[comparePos], numList[i] = numList[i], numList[comparePos]
			// 修改比较位
			comparePos = i

			fmt.Println("交换后: ", numList)

			//lower := numList[i]
			// 先将这个数 从切片中删除
			//numList = append(numList[:i], numList[i+1:]...)
			// 再插入到第一个位置

		} else if i < comparePos && numList[i] > compareV {
			fmt.Println("[", i ,"]:", numList[i],
				"compareV[:",comparePos,"]", compareV)

			// 和比较位交换
			numList[comparePos], numList[i] = numList[i], numList[comparePos]
			// 修改比较位
			comparePos = i
			fmt.Println(numList)
		}
	}
	fmt.Println("end", numList, "numList[",comparePos,"]",compareV)

	return comparePos
}

// 错误版本1
func test1(numList []int, lower int, high int, count *int) int {
	*count = *count + 1
	fmt.Println("----------- ", *count," -----------")

	// 步骤1
	// 将 3 1 4 2 5  固定3  1 2 3 4 5
	// 2 3 1 4 5

	// 比较位
	var compareV = numList[lower]
	var comparePos = lower

	fmt.Println("start", numList,
		"comparePos: ", comparePos, " compareV: ", compareV)

	for i := high; i > lower; i-- {
		// 小于放左端
		if i > comparePos && numList[i] <= compareV {

			fmt.Println("compareV[",comparePos, "]:", compareV,
			"[", i , "]: ", numList[i],)

			// 和比较位交换
			numList[comparePos], numList[i] = numList[i], numList[comparePos]
			// 修改比较位
			comparePos = i

			fmt.Println("交换后: ", numList)

			//lower := numList[i]
			// 先将这个数 从切片中删除
			//numList = append(numList[:i], numList[i+1:]...)
			// 再插入到第一个位置

		} else if i < comparePos && numList[i] > compareV {
			fmt.Println("[", i ,"]:", numList[i],
				"compareV[:",comparePos,"]", compareV)

			// 和比较位交换
			numList[comparePos], numList[i] = numList[i], numList[comparePos]
			// 修改比较位
			comparePos = i
			fmt.Println(numList)
		}
	}
	fmt.Println("end", numList, "numList[",comparePos,"]",compareV)
	//*comparePos1 = comparePos

	return comparePos

	//result[comparePos] = compareV
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

