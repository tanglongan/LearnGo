package main

import (
	"fmt"
	"strings"
)

/*
 * 闭包
 * 闭包是一个函数，这个函数包含了它的外部作用域中的变量
 * 底层原理
 * 1.函数可以作为返回值
 * 2.函数内部查找变量的顺序，现在自己内部找，找不到往外层找
 */
func main() {
	ret := addr()
	ret2 := ret(200)
	fmt.Println(ret2)

	fmt.Println("===================")

	rr := f3(f2, 500, 300) //把原来需要传递两个int类型的参数包装成一个不需要传参的函数
	f1(rr)

	fmt.Println("===================")
	testMakeSuffixFunc()
}

func addr() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

//有两个函数f1、f2。要求实现f1(f2)
func f1(f func()) {
	fmt.Println("this is f1")
	f()
}

func f2(x, y int) {
	fmt.Println("this is f2")
	fmt.Println(x + y)
}

func f3(f func(int, int), x, y int) func() {
	temp := func() {
		f(x, y)
	}
	return temp
}

//示例
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		//如果name串不是以suffix结尾的为，就拼接在一起
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func testMakeSuffixFunc() {
	jpgFunc := makeSuffixFunc(".jpg")
	txtFunc := makeSuffixFunc(".txt")

	fmt.Println(jpgFunc("test"))
	fmt.Println(txtFunc("hehe"))
}
