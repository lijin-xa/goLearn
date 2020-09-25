/*
* 入口函数 main
* 同路径下 文件 - packageName - 必须一致
* package 命名规则
* 1 包名一般小写（简单且有意义）
* 2 包名一般和所在目录名相同，包名不能包含_符号
* 3 一个源文件夹下的所有源码文件只能属于同一个包
* 调用其他包的方法 - packageName.funcName
* 单个文字的方法如果想外部调用 首字母必须大写（相当于publick） 内部使用（private）
* 单个文件下 运行时会自动调用 init()函数  可以有同名的init函数 会以此按顺序调用
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"test/function"
	"test/goroutine"
	_interface "test/interface"
	_package "test/package"
	_redis "test/redis"
	_struct "test/struct"
)

//func init() {
//	fmt.Println("sorry my first come out")
//}

func testRedis() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "192.168.196.50:6007",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping(context.Background()).Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}

// 主入口函数
func main() {
	// 调用mongodb包
	//mongodb.Start()
	// 调用redis包
	_redis.Start()

	//callGoroutine()
	//callPackage()
	//callInterface()
	//callStruct()
}

// 调用function包
func callFunction() {
	function.Start()
	function.StartInterface()
	function.StartClosure()
	function.StartPanic()
	function.StartReCover()
}

// 调用 _struct包
func callStruct() {
	//_struct.Start()
	_struct.StartOOP()
}

// 调用 interface包
func callInterface() {
	//_interface.Start()
	_interface.StartSort()
}

// 调用 package包
func callPackage() {
	//_package.StartTime()
	//_package.StartOs()
	_package.StartFlag()
}

// 调用goroutine包
func callGoroutine() {
	goroutine.Start()
}