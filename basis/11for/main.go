package main

import "fmt"

//for循环：go语言中只有一种循环for
func main() {
	testFor1()
	testFor99()
}

//for循环的基本形式与变种
func testFor1() {
	//基本格式
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}

	//变种2
	var j = 20
	for j < 22 {
		fmt.Println(j)
		j++
	}

	//变种3：for range循环
	s := "hello World"
	for index, value := range s {
		fmt.Printf("%d %c\n", index, value)
	}
}

//打印九九乘法表
func testFor99() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
