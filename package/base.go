package _package
// 包的基础概念

/*
// 包的导入路径 GOPATH/src/lab/a
绝对路径 import "lab/a"
相对路径 import "../a"

// 包的引用格式
1 标准格式 import "fmt"
2 自定义别名 import F "fmt" F就是包的别名  - 可以避免导入相同的包命名冲突
3 省略引用格式 import . "fmt"  调用包方法可以省略包名
4 匿名引用格式 import _ "fmt"  只希望执行包的初始化init函数，而不使用包内部数据

// 引用包时需要注意
1 一个包可以有多个init函数，包加载的时候会全部执行，但是不能保证执行顺序
2 包不能出现环行引用 允许重复引用

// 包的初始化
1 从main函数引用的包可以，逐级查找包的引用
2 单个包的初始化 先初始化常量 - 全局变量 - init函数

// $GoROOT/src/pkg - Go语言代码库中的包 例如: fmt io
// 常用的包 - 标准库
fmt 格式化的标准数据输出

io 原始的I/O操作界面 对os包这样的原始I/O进行封装

bufio 通过对io包的封装 提供了数据缓冲功能

sort 对切片和用户定义的集合进行排序的功能

strconv 字符串和基本类型之间的相互转换

os 不依赖平台的操作系统函数接口

sync 实现多线程中锁机制以及其他同步互斥机制

flag 提供命令行参数的规则定义和传入参数解析的功能

encoding/json 提供了对json的基本支持

html/template web开发中生成html的 template 的一些函数

net/http 提供了http相关服务 - http请求、响应和URL解析

reflect 实现了运行时反射 允许程序通过抽象类型操作对象

os/exec 执行自定义linux命令的相关实现

strings 处理字符串的一些函数集合

bytes 对字节切片进行读写操作的一系列函数

log 在程序中输出日志


// import 导入包时，使用的包所属文件夹的名字
// 导出包中的标识符 - 让外部访问包的类型和值 则需要将访问数据的首字母大写

// init 函数特性
1 每个源码可以使用1个init()函数
2 init()在main()函数执行前被调用
3 调用顺序为main()中引用包，以深度优先顺序初始化 最后引用的最先调用
*/

// big包 - 对整数的高精度计算
// regexp - 正则表达式

















