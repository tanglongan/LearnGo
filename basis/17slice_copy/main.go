package main

import "fmt"

// 切片copy
func main() {
	a1 := []int{1, 2, 3}
	a2 := a1
	var a3 = make([]int, 3) //定义变量，同时开辟内存空间
	copy(a3, a1)            // copy
	fmt.Println(a1, a2, a3)

	a1[0] = 100
	fmt.Println(a1, a2, a3)

	var a4 []int //只是定义了变量，没有开辟内存空间，因此copy没生效
	copy(a4, a1)
	fmt.Println(a4)

	//创建切片，长度为5，容量为10
	var a = make([]int, 5, 10)
	for i := 1; i <= 10; i++ {
		a = append(a, i)
	}
	fmt.Println(a)
	fmt.Println(cap(a))
}
