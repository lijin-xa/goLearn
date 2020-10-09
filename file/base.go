package file

import (
	"bufio"
	"encoding/gob"
	"encoding/xml"
	"fmt"
	"os"
)

// 文件处理
// 提前了解一下 go - json文件的读写操作 内置库在encoding/json
// ` 不是单引号 标准语法叫做标签Tag
type User struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
}

func Start() {
	// 返回一个json encoding
	//user := &User{UserId: 1, UserName: "lijin"}
	//j, _ := json.Marshal(user)
	//fmt.Println(string(j))

	//testXML()
	//testGob()
	testTxt()
}

type Website struct {
	Name   string `xml:"name,attr"`
	Url    string
	Course []string
}

// 写入XML文件
func testXML() {
	info := Website{"go", "www.go.com", []string{"go-1", "go-2"}}
	f, err := os.Create("./info.xml")
	if err != nil {
		fmt.Printf("create file failed, err %s\n", err)
		return
	}
	defer f.Close()

	encoder := xml.NewEncoder(f)
	err = encoder.Encode(info)
	if err != nil {
		fmt.Println("encode failed")
		return
	} else {
		fmt.Println("encode success")
		fmt.Println(info)
	}

	openXML()
}

func openXML() {
	// 打开文件
	f, _ := os.Open("./info.xml")

	defer f.Close()

	info := &Website{}

	// Decode解码 Encode编码
	decoder := xml.NewDecoder(f)
	decoder.Decode(info)
	fmt.Println("decode: ", info)
}

// 练习Gob
func testGob() {
	//
	//name := "demo.gob"
	//
	//info := map[string]string {
	//	"name": "go",
	//	"url": "www.go.com",
	//}
	//name := "./demo.gob"
	//
	//// | 位运算符 按位或
	//f, _ := os.OpenFile(name, os.O_RDWR|os.O_CREATE, 0777)
	//defer f.Close()
	//
	//encode := gob.NewEncoder(f)
	//if err := encode.Encode(info); err != nil {
	//	fmt.Printf("gob encode failed, err%s\n", err)
	//	return
	//}
	openGob()
}

// 解码Gob文件
func openGob() {
	var m map[string]string
	f, _ := os.Open("demo.gob")
	d := gob.NewDecoder(f)
	d.Decode(&m)
	fmt.Println(m)
}

// txt文件
func testTxt() {
	//filePath := "./test.txt"
	//file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	fmt.Printf("openfile failed, err %s\n", err)
	//	return
	//}
	//defer file.Close()
	//
	//str := "this is a go file \n"
	//writer := bufio.NewWriter(file)
	//writer.WriteString(str)
	//writer.Flush()

	openTxt()
}

func openTxt() {
	// 打开文件
	file, err := os.Open("./test.txt")
	if err != nil {
		fmt.Printf("open file failed, err %s\n", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n')
	fmt.Printf("test.txt content: %s\n", str)
	fmt.Println("txt read finish")
}