/*
* list 列表 底层实现是由双向链表
* 没有具体元素类型的限制
* 如果给列表中放入了一个 interface{} 类型的值
* 取出值后，如果要将 interface{} 转换为其他类型将会发生宕机
*/

package main

import (
	"container/list"
	"fmt"
)

func main() {

	/*
	* 初始化 list
	* 1 变量名 := list.New()
	* 2 var 变量名 list.List
	*/
	list1 := list.New()

	// 尾部插入
	list1.PushBack( "go")

	// 头部插入
	list1.PushFront("container")

	// 返回的是一个添加的元素句柄
	element := list1.PushBack("list")

	// 在list 之后添加
	list1.InsertAfter("and", element)

	// 在list 之前添加
	list1.InsertBefore("or", element)

	// 遍历list
	for i := list1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}


