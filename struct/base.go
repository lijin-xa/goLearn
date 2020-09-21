// Go 语言中没有 类 的概念 也不支持类的继承等面向对象的概念
// Go 结构体内嵌配合接口比面向对象具备更高的扩展性和灵活性=
package _struct

import "fmt"

// User 结构体名 （注意这里没有分割符号） 但是对结构体字段初始化时需要加上,分隔符
type User struct {
	uid int
	name string
	age int
}

func Start() {
	// 使用new关键字对结构体进行初始化
	user := new(User)

	// 与C/C++不同的是 结构体指针指向结构体内的成员变量时不再使用 ->
	// 而是统一使用 . 其实是使用了一个语法糖user.name == (*user).name
	user.uid = 100001
	user.name = "Tim"
	user.age = 18

	// 取结构体的地址实例化 实际上对结构体进行了一次new的实例化
	user1 := &User{}
	user1.uid = 100002
	user1.name = "jack"
	user1.age = 19

	user2 := newUser(100003,"Lisa", 20)
	fmt.Println(user2)

	// 匿名结构体实现
	anonymousStruct()
}

// 取地址实例化是最广泛的一种结构体实例化的方式，可以使用函数来封装初始化的过程
func newUser(uid int, name string, age int) *User {
	return &User {
		uid: uid,
		name: name,
		age: age,
	}
}

func printMsgType(msg *struct{
	id int
	data string
}) {

	fmt.Printf("%+v", msg)
}

// 匿名结构体操作
func anonymousStruct() {
	msg := &struct {
		id int
		data string
	}{
		1,
		"go",
	}

	printMsgType(msg)
}