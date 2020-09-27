package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 针对第一个问题做回溯
// 1 mongodb 操作
// 2 接口如何使用
// 3 切片的创建
// 4 随机函数
// 5 排序 - QuickSort()

// 声明一个接口
type name interface {
	String()
}

// 用于实现接口
type jam struct {
}

func (j *jam) String() {
	fmt.Println("It is implement interface - String function.")
}

// 接口
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

	// 插入100个随机元素
	for i := 0; i < length; i++ {
		numList[i] = rand.Intn(length)
	}

	// 随机函数 rand.Intn(n int64))
	// 是一个伪随机函数 不管运行多少次 返回的随机数都是固定的
	// 如果想变化 必须传入随机种子 rand.Seed(seed int64)
	// 目前了解到的 随机种子传入纳秒时间戳
	// rand.Seed(time.Now().UnixNano())

	//fmt.Println(len(numList))
	return numList
}

func main() {
	// 实现一个接口
	//interfaceDemo()
	fmt.Println("时间: ", time.Now().UnixNano())

	// 创建一个长度为100的切片 并且随机填充1~100的数
	numList := sliceDemo(100)
	//var numList = []int{1, 7, 7, 9, 1, 8, 5, 0, 6, 0}
	lower := 0
	high := len(numList) - 1
	var a = 0
	count := &a // 测试打印日志用

	fmt.Println("sort start: ", numList)
	quickSort(numList, lower, high, count)
	fmt.Println("sort end: ", numList)
}

// 快速排序
func quickSort(numList []int, lower int, high int, count *int) {
	comparePos, numList := stepSort(numList, lower, high, count)
	//fmt.Println("the step numList:", numList,
	//	" comparePos:",  comparePos)

	// 左端排序
	if lower < comparePos {
		//fmt.Println("lower: ", lower, " high: ", comparePos-1)
		//fmt.Println("左端数: ", numList[lower:comparePos])
		quickSort(numList, lower, comparePos-1, count)
	}
	// 右端排序
	if high > comparePos {
		//fmt.Println("high: ", high, " lower: ", comparePos+1)
		//fmt.Println("右端数: ", numList[comparePos+1:])
		quickSort(numList, comparePos+1, high, count)
	}
}

// 步骤1排序 实现 在指定区域内 比较位左端的数均小于它 右端的数均大于它
func stepSort(numList []int, lower int, high int, count *int) (int, []int) {
	*count = *count + 1
	//fmt.Println("----------- ", *count," -----------")
	// 比较位
	var compareV = numList[lower]
	var comparePos = lower

	//fmt.Println("start", numList,
	//	"comparePos: ", comparePos, " compareV: ", compareV)

	// 暂时不考虑 性能问题
	// 对指定区域数据移动位置 将小于比较位的数据 均移动到比较位左边
	// 简单的设计 将这个数移动区域的第一个位置
	// 从高位依次移动
	move := high
	for i := high; i > lower; i-- {
		if numList[move] <= compareV {
			moveV := numList[move]

			// TODO 在这里存在严重的性能问题 - 日后需要进行优化
			// 将该位置的元素从切片中移出
			numList = append(numList[:move], numList[move+1:]...)

			// 添加到切片的首位置
			// 修改为插入指定区域的首位置
			//numList = append([]int{moveV}, numList...)
			numList = append(numList[:lower], append([]int{moveV}, numList[lower:]...)...)
			comparePos = comparePos + 1 // 比较位后移一位

			//fmt.Println("move", move, " - ", numList)
		} else {
			move = move - 1
		}
	}
	//fmt.Println("end", numList, "numList[",comparePos,"]:",compareV)
	return comparePos, numList
}

// 错误版本1（留作记录）
func test1(numList []int, lower int, high int, count *int) int {
	*count = *count + 1
	fmt.Println("----------- ", *count, " -----------")

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

			fmt.Println("compareV[", comparePos, "]:", compareV,
				"[", i, "]: ", numList[i])

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
			fmt.Println("[", i, "]:", numList[i],
				"compareV[:", comparePos, "]", compareV)

			// 和比较位交换
			numList[comparePos], numList[i] = numList[i], numList[comparePos]
			// 修改比较位
			comparePos = i
			fmt.Println(numList)
		}
	}
	fmt.Println("end", numList, "numList[", comparePos, "]", compareV)
	//*comparePos1 = comparePos

	return comparePos
}

// 练习代码
func sliceProject() {

	// 切片的创建方式 切片的数组的区别
	//var sli1 = []int{1, 7, 7, 9, 1, 8, 5, 0, 6, 0}

	// make 动态创建一个切片 size=10 capacity=20
	// 设置size之后，会生成默认数据
	//sli2 := make([]int, 0, 20)

	// 往切片中追加
	//sli1 = append(sli1, 1)
	//fmt.Println(sli1)

	//sli2 = append(sli2, 1)
	//fmt.Println(sli2)

	// 期望的输出结果    1 0 0 1 7 7 9 8 5 6  sli1[3]=1
	//test2(sli1, 0, 9)

	//numList := sli1
	//lower := 0
	//high := len(numList) - 1
	//var a = 0
	//count := &a
	//quickSort(numList, lower, high, count)
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
