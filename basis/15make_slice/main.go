package main

import "fmt"

/*
 * make函数创造切片
 * 切片的本质就是一个框，框住了一块连续的内存
 * 切片是一个引用类型
 * 切片之间是不能比较的，不能使用==操作符来判断两个切片是够含有全部相等的元素。
 * 切片唯一合法的比较操作时和nil比较，一个nil值的切片并没有底层数组，一个nil值的切片的长度和容量都是0。但是不能说一个长度和容量都是0的切片一定是nil
 **/
func main() {
	// 切片类型、长度、容量
	s1 := make([]int, 5, 10)
	fmt.Printf("s1=%v len(s1)=%d cap(s1)=%d\n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2=%v len(s2)=%d cap(s1)=%d\n", s2, len(s2), cap(s2))

	fmt.Println("========================")
	s3 := []int{1, 3, 5}
	s4 := s3 //s3和s4指向了同一个底层数组
	fmt.Println(s3, s4)

	s3[0] = 1000 //修改切片元素值
	fmt.Println(s3, s4)

	fmt.Println("========================")
	//遍历方式一：索引
	for i := 0; i < len(s3); i++ {
		fmt.Println(s3[i])
	}
	fmt.Println("========================")
	//遍历方式二：range循环
	for _, v := range s3 {
		fmt.Println(v)
	}
}
