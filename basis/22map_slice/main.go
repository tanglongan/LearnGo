package main

import "fmt"

//map和slice组合
func main() {

	//元素类型为map的切片
	var s1 = make([]map[int]string, 10)
	s1[0] = make(map[int]string, 1) //对内部的map做初始化
	s1[0][10] = "北京"
	fmt.Println(s1)

	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30} //对内部的切片进行初始化
	fmt.Println(m1)

}
