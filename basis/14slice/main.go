package main

import "fmt"

/*
 * Go语言的切片
 * 切片是一个拥有想同类型元素的可变长度的序列。它是基于数组类型做的一层封装。非常灵活，支持自动扩容
 * 切片是一个引用类型，它的内部结构包含地址、长度、和容量。
 * 切片一般用于快速地操作一块数据集合
 * 切片定义：var name []T  name变量名，T表示切片中的元素类型
 * 切片拥有自己的长度和容量，可以日通过内置的len()函数求长度，使用内置的cap()函数求切片的容量
 * 基于数组定义切片：由于切片的底层就是一个数组，所以可以基于数组定义切片
 */
func main() {
	// 1.直接定义切片
	var a []string              //声明一个字符串切片并初始化
	var b = []int{}             //声明一个整型切片并初始化
	var c = []bool{false, true} //声明一个布尔切片并初始化

	a = []string{"北京", "上海", "深圳", "广州"}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// 2.通过切片表达式得到切片。切片表达式：一种指定low和high两个索引界限值的简单形式，另一种是除了low和high索引界限值外还指定了容量的完整的形式
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s1 := a1[0:4] //基于一个数组切割，左包含，右不包含
	fmt.Println(s1)
	fmt.Println(a1[1:6])
	fmt.Println(a1[:4])

	//切片长度：元素的个数
	//切片容量：底层数组从切片的第一个元素到最后一个元素的数量
	s2 := a1[:4]
	fmt.Printf("len(s2):%d cap(s2):%d\n", len(s2), cap(s2))

	//切片再切片
	s3 := a1[3:]
	s4 := s3[1:]
	fmt.Printf("len(s4):%d cap(s4):%d\n", len(s4), cap(s4))

}
