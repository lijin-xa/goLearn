package must

import (
	"fmt"
)

func test() {

	var sli = make([]int, 100)
	for i := 0; i < 100; i++ {

		//append(sli, r)
	}
	for k, v := range sli {
		fmt.Println(k, v)
	}

	//fmt.Println("len(sli)", len(sli))

	//sli = sli()
	len := len(sli)-1
	for i := 1; i <= len; i++ {
		for j := 1; j <= len - i; j++ {
			if sli[j] > sli[j+1] {
				sli[j], sli[j+1] = sli[j+1], sli[j]
			}
		}
	}

	for k, v := range sli {
		fmt.Println(k, v)
	}
}

// TODO
// mongodb 操作
// 接口如何使用
// 切片的创建
// 随机函数



