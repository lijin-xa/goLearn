/*
* map映射 和STL map容器类似
* pair<key, value> go 无须集合 - STL类模板
* 声明一个map var mapName map[keyType]valueType
*/

package container

import (
	"fmt"
	"sync"
)

// 入口程序
func StartMap(){
	// 声明一个map key - string value - int
	// var mapList map[string]int

	// make函数 创建一个指定大小的map
	mapList := make(map[string]int)
	mapList["key1"] = 1
	mapList["key2"] = 2

	// 遍历输出 无法得到指定的顺序输出
	for k, v := range mapList {
		fmt.Printf("key: %s, value: %d\n", k, v)
	}

	// 删除map指定键值行 delete(map, mapKey)
	delete(mapList, "key1")
	for k, v := range mapList {
		fmt.Println(k, v)
	}


	/*
	* go并没有提高清空map的接口， 直接使用make重新建立一个新map
	* go并行垃圾回收的效率比写一个清空map函数效率要高的多
	* 多键值索引 目前只是做大概阅读
	*/

	/*
	实例化一个map 键类型int 值类型*Profile
	自主生成哈希值 来实现多键值查询
	type Profile struct {
		Name string
		Age int
		Married bool
	}
	var mapper = make(map[int][] *Profile)
	var mapper = make(map[int][] *Profile)
	*/

	/*
	* map 在并发情况下 同时读写是线程不安全的， 一般处理并发比如C++进行加锁 mutex，Go1.9版本中加入了sync.Map包 支持并发安全
	* 1 无须初始化 直接声明
	* 2 去除原本的map方式进行取值和设置，改用包内方法 Store - 存储 Load - 获取 Delete - 删除
	* 3 使用Range配合回调函数进行遍历 回调函数需要继续遍历返回true 终止返回false
	*/
	var scene sync.Map
	scene.Store("red", 1)
	scene.Store("blue", 2)
	scene.Store("yellow", 3)

	// 返回value error
	fmt.Println(scene.Load("red"))

	// 删除指定键值的数据行
	scene.Delete("blue")

	// 遍历sync.Map中的键值对 interface{}类似一个 void*指针 任意类型
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterator: ", k, v)
		return true
	})
}
