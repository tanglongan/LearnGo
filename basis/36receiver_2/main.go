package main

import "fmt"

/*
 * 为任意类型添加方法
 */

//定义一个新类型
type MyInt int

//为MyInt类型添加一个SayHello的方法
func (m MyInt) SayHello() {
	fmt.Println("Hello, 我是一个Int")
}

func main() {
	var m1 MyInt
	m1.SayHello()
	m1 = 100
	fmt.Printf("%#v\n%T\n", m1, m1)
}
