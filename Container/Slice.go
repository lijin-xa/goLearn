/*
* 切片 Slice 基础概念
*/
package main

import "fmt"

func main(){

	var str string = "hello use Printf"
	// %T输出 数据类型 %v输出值
	fmt.Printf("%T, %v\n", str, str)

	/**
	* 切片复制 copy()
	* 签名 copy(destSlice, srcSlice[]T) int 将srcSlice复制到desSlice
	* @param destSlice 复制的目标
	* @param srcSlice 数据来源切片
	* 返回实际发生复制的元素个数
	*/
	var str1 []string = []string{"_", "implement", "copy"}
	str2 := []string{"slice"}

	// str2 复制到 str1的初始位置 具体复制情况必须得基于双方的空间大小
	copy(str1, str2)
	fmt.Println(str1)

	const elementNum = 100
	srcData := make([]int, elementNum)

	for i := 0; i<elementNum; i++ {
		srcData[i] = i
	}

	// srcData 引用数据
	refData := srcData

	destData := make([]int, elementNum)
	copy(destData, srcData)

	// 会影响refData 但不会影响destData
	srcData[0] = 101

	fmt.Println(refData[0])
	fmt.Println(destData[0])
	/*
	* 切片删除元素
	* a[1:2] 和 a[1:]
	* 如果结束位置缺省，则可以输出从起始位置起的全部切片元素
	* 如果没有省略 则只能输出从起始位置 到 结束位置-1的元素
	*/
	a := []int{1, 2, 3}

	// 删除开头一个元素
	// 1 相当于移动指针的位置从 0 - 1
	// a = a[1:]

	// 2 将后面的元素向开头移动 相当于将a[1:]位置的元素移动到0位置
	// a = append(a[:0], a[1:]...)

	/*
	* copy(a, a[1:]) 相当于是将 a[1:] 1 - 2位置的元素插入到a 0 - 1位置 （copy不能超出a原范围）
	* copy函数 返回时发生复制的元素个数
	* a[:copy(a, a[1:])] == a[:2]
	*/
	// 3 删除开头1/N个元素
	// a = a[:copy(a, a[1:])]
	// 从尾部删除1个元素 a[0:2] 只会输出1 2
	// a = a[:len(a)-1]

	/*
	* 删除指定位置index的元素
	* arr = append(arr[:index], arr[index+1:]...)
	* 将arr切片index+1位置后的元素 全部追加到index位置之后
	*/
	// 删除1位置的元素
	a = append(a[:1], a[2:]...)

	// range 关键字 遍历切片
	slice1 := []int{1, 2, 3, 4}
	for k, v := range slice1 {
		fmt.Printf("%T, %d\n", k, v)
	}

	// 多维切片 简单理解基础概念
	// 声明并赋值一个二维切片
	slice2 := [][]int{{1}, {10, 20}}

	// 往二维切片的内第一个切片插入一个数据
	slice2[0] = append(slice2[0], 2)
	for k, _ := range slice2 {
		for k1, v1 := range slice2[k]{
			fmt.Printf("二维切片 %T, %d\n", k1, v1)
		}
	}
}