package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
)

var rdb *redis.Client
func initClient() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

// 问题1 Mysql 查询表中所有记录
// 使用 mysql函数 count - Mysql函数这块不是很熟悉
// sql语句 select count(*) from table_name
// count(*) 表示所有字段
// count(field) 查找单字段
// mongo 查询集合中文档的数量
// 使用游标 db.collectionName.find().count()

// 问题2 go-driver 操作redis
// 只是大概了解到具体流程 直接手写记忆不熟

// 问题3 快速排序 不熟练 - 花费大量时间
// 在步骤1 - 实现二分法   缺乏对快排算法思维的理解
func main() {
	////var sli1 = []int{}
	sli1 := make([]int, 1000)
	for i := 0; i < 1000; i++ {
		r := rand.Intn(1000)
		sli1[i] = r
	}
	//sli1 := []int{5, 3, 4, 2, 1, 1,2,6, 8, 7}
	a := 0
	count := &a
	fmt.Println("sort start: ", sli1)
	quickSort1(sli1, 0, len(sli1)-1, count)
	fmt.Println("sort end: ", sli1)
}

// 快速排序（第二版）
func quickSort1(numList []int, low int, high int, count *int) {
	comparePos := step2(numList, low, high, count)

	if comparePos > low {
		//fmt.Println("left sort: ", numList[low:comparePos])
		quickSort1(numList, low, comparePos-1, count)
	}
	if comparePos < high {
		//fmt.Println("right sort: ", numList[comparePos+1:])
		quickSort1(numList, comparePos+1, high, count)
	}
}

// 优化前 - 老版本（使用切片的元素移位 元素位置的移动）
func step1(numList []int, low int, high int, count *int) (int, []int) {
	// 比较位
	compareV := numList[low]
	comparePos := low

	//fmt.Println("numList: ", numList, "low: ", low, "high: ", high, "compareV: ", compareV)

	// 实现让比较位 左边均小于它 右边均大于它
	// 3 2 1 4 5 --- 2 1 3 4 5
	movePos := high
	for i := high; i > low; i-- {
		if numList[movePos] <= compareV {

			// 算法设计 - 需要注意的地方
			// 将移动元素移出 后续元素全部向前移动一位
			// 再将该元素插入到当前的比较位的位置
			// 所以虽然循环在继续 但是实际上下一次比较的位置仍然在上一次的原位置
			move := numList[movePos] // 移动位置的元素

			// 1 将该元素从切片中移出
			numList = append(numList[:movePos], numList[movePos+1:]...)
			//fmt.Println("1 将该元素从切片中移出: ", numList)

			// 2 将移出元素插入指定区域的首位置
			// 3 将比较位置后的所有元素都追加到move元素后面
			// 4 将3步骤组合而成的元素追加到比较位置前的元素
			// 之所以切片位置[:comparePos] 选择comparePos 没有-1 是因为指定切片的结束位置 则不包括该位置的数
			// append([]int{move}, numList[comparePos:]...)

			// 这里的插入位置可以选择 初始的指定区域的首位置 也可以选择向前移动后的比较位
			numList = append(numList[:comparePos], append([]int{move}, numList[comparePos:]...)...)
			comparePos = comparePos + 1

			//fmt.Println("numList: ", numList)
		} else {
			movePos = movePos - 1
		}
	}

	*count = *count + 1
	//fmt.Println("step1: ",*count, "numList: ", numList, " sort pos:",comparePos,"value: ", compareV)

	// 之所以再返回 - 追加之后numList相当于重新复制了一次
	return comparePos, numList
}

// 优化上次快速排序未完的算法
// 目标 - 将指定范围的通过一定的移动操作 实现比较位左边均小于它 右边均大于它
// 默认比较数为范围第一个数
// 1 将比较位独立出来， 将剩余区域内元素 在指定的高位坐标 低位坐标+1
// 同时移动两个坐标 高位坐标-1 低位坐标+1 直至两个坐标相遇结束
// 2 如果在移动过程中 出现高位坐标位置的数 <= 比较数
// 低位坐标位置的数 > 比较数 将这两个数交换位置
// 实现 - 两个坐标相遇的位置 左边的数均 <= 比较数 右边的数 > 比较数
// 3 将比较数插入到这个位置 - 即将该位置数同 比较数的位置交换
func step2(number []int, low int, high int, count *int) int {
	// 由于go没有while语句 这里使用for{} 来代替
	compareV := number[low]
	tempLow := low + 1
	for {
		// 高端坐标向低端位置移动 - 直到出现 <= 比较数
		// 高端位置作为最后比较位的确定位置  最多可以移动到比较位的位置
		for ;high >= tempLow; {
			if number[high] <= compareV {
				break
			} else {
				high--
			}
		}

		// 低端坐标向高端位置移动 - 直到出现 > 比较数
		for ;tempLow < high; {
			// 先判断当前位置是否符合
			if number[tempLow] > compareV {
				break
			} else {
				tempLow++
			}
		}

		// 高低端坐标相遇 结束
		if tempLow >= high {
			break
		} else {  // 交换两个数据的元素
			number[tempLow], number[high] = number[high], number[tempLow]
		}
	}

	// 将比较数插入的两端坐标交汇的地方
	number[low], number[high] = number[high], number[low]
	//fmt.Println("step number: ", number)

	*count = *count + 1
	//fmt.Println("step1: ",*count, "numList: ", number, " sort pos:",high,"value: ", compareV)
	return high
}
