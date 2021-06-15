package main

import "fmt"

func main() {
	testSliceAndArray()
	fmt.Println("=============")
	testMake()
	fmt.Println("=============")
	testAppend()
	fmt.Println("=============")
	testRmLast()
}

/*
 * 比较切片和数组的区别
 * 1.数组是值类型，将一个数组复制给另一个数组，传递的是一份拷贝
 * 2.切片是引用类型，切片包装的数组称为切片的底层数组
 **/
func testSliceAndArray() {
	//a是一个数组，注意数组是一个固定长度的，初始化时必须要指定长度，不指定长度的话就是切片了
	a := [3]int{1, 2, 3}
	b := a    //b是数组，是a数组的一份拷贝
	c := a[:] //c是切片，是引用类型，底层数组是a

	for i := 0; i < len(a); i++ {
		a[i] = a[i] + 1
	}

	//改变了a数组中的值之后，b是a的拷贝，因此b不变，c是引用，c的值会改变
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}

/*
 * make只能用于slice、map和channel，所以下面函数生成的slice是一个引用类型
 */
func testMake() {
	s1 := make([]int, 0, 3)
	for i := 0; i < cap(s1); i++ {
		s1 = append(s1, i)
	}

	s2 := s1

	for i := 0; i < len(s2); i++ {
		s2[i] = s2[i] + 1
	}

	fmt.Printf("s1: %v\n", s1)
	fmt.Printf("s2: %v\n", s2)
}

/*
 * 当对slice append超出底层数组的界限时，会产生一个新的底层数组
 */
func testAppend() {
	n1 := [3]int{1, 2, 3}
	n2 := n1[0:3]
	fmt.Println("address of items in n1:")
	for i := 0; i < len(n1); i++ {
		fmt.Printf("%p\n", &n1[i])
	}

	fmt.Println("address of items in n2:")
	for i := 0; i < len(n2); i++ {
		fmt.Printf("%p\n", &n2[i])
	}

	//对n2执行append操作之后，n2超出了底层数组n1的j
	n2 = append(n2, 5)
	fmt.Println("address of items in n1:")
	for i := 0; i < len(n1); i++ {
		fmt.Printf("%p\n", &n1[i])
	}

	fmt.Println("address of items in n2:")
	for i := 0; i < len(n2); i++ {
		fmt.Printf("%p\n", &n2[i])
	}
}

/*
 * 实现了删除slice最后一个item的函数
 */
func rmLast(a []int) {
	fmt.Printf("[rmLast] the address of a is %p\n", a)
	a = a[:len(a)-1]
	fmt.Printf("[rmLast] after remove, the address of a is %p\n", a)
}

func testRmLast() {
	xyz := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("[main] the address of xyz is %p\n", xyz)

	rmLast(xyz)

	fmt.Printf("[main] after remove, the address of xyz is %p\n", xyz)
	fmt.Printf("%v", xyz)
}
