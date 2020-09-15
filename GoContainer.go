/**
* go container 语言容器
*
*/

package main

import "fmt"

func main(){

	/**
	* arr arrayValue
	* [3]int{1, 2, 3} [3]数组长度
	*/
	//var arr [3]int = [3]int{1, 2, 3}
	//var arr1[3]int

	//[...] 数组的长度根据初始化来计算
	//arr1 = [...]int{4, 5, 6}
	//for k, v := range arr{
	//	fmt.Printf("%T, %v \n", k, v)
	//	fmt.Println("  \n", k, v)
	//}
	//
	//for k, v := range arr1{
	//	fmt.Println("  \n", k, v)
	//}


	/**
	* 多维数组 数组的下标是从0开始的
	*/
	var multiArr [2][4]int = [2][4]int{{1, 2, 3, 4},{11, 22, 33, 44}}
	fmt.Println(multiArr[1][3])


	var tempArr = [3]int{1, 2, 3}

	/**
	* slice[startPos:endPos]
	* tempArr[1:2]生成新切片  输出[1 2]
	* 取出元素不包含结束为止对应的索引  使用slice[len(slice)]取出
	* 缺省startPos 默认连续区域起始位置
	* 缺省endPos 默认连续区域结束位置
	* len(arr) 数组长度
	*/
	fmt.Println(tempArr)
	fmt.Println(tempArr[0:2])
	tempArr1 := tempArr[0:2]

	//和原切片一致
	tempArr2 := tempArr[:]

	//重置切片 输出[]
	tempArr3 := tempArr[0:0]
	fmt.Println(len(tempArr1))
	fmt.Println(tempArr2)
	fmt.Println(tempArr3)


	/**
	* 直接声明切片 var sliceName []sliceType
	* 可以使用append函数 往切片中添加元素
	* 切片的扩容跟STL vector的扩容机制基本一致 vector linux下扩容机制也是以2的倍数进行扩容
 	*/
	var strList []string
	fmt.Println(strList == nil)
	strList = append(strList, "go")
	strList = append(strList, "container")

	//往切片的头部插入数据 但效率很差 会将已有元素再复制一次
	strList = append([]string{"no"}, strList...)

	/**
	* 往切片的指定位置插入数据 （个人感觉挺麻烦的）
	* 插入时会先复制原有元素 no go container 再往no - go 中间插入off
	* 结果[no off no go container]
	*/
	strList = append(strList[:1], append([]string{"off"}, strList...)...)
	fmt.Println(strList)

	var numList []int
	for i :=0; i < 10; i++ {
		numList = append(numList, i)
	}
	//cap() 查看切片容量
	fmt.Printf("len: %d cap: %d\n", len(numList), cap(numList))



	/**
	* make()函数构造切片  make([]type, size, cap)
	* 使用make函数生成的切片会发生内存分配的操作
	* @param type 切片的元素类型
	* @param size 分配了多少元素
	* @param cap 切片的容量
	*/
	sli1 := make([]int, 2, 4)

	//输出[0 0]
	fmt.Println(sli1)
}