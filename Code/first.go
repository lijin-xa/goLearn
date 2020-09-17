/*
* 第一个简单go程序
* 1 go mod init xx 适用7牛中国区代理 下载go依赖包 https://goproxy.cn,direct
* File - Settings - Go - Go Modules 引用GOPROXY=https://goproxy.cn,direct
* 来实现不用在指定目录下src编写项目
* func 函数 import 外部引用
* fmt.Println 输出
*/
package code

import "fmt"

func SimpleGo(){
	fmt.Println("hello world")
}



