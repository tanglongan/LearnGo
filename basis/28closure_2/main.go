package main

import "fmt"

func main() {
	f1, f2 := calc(10)

	fmt.Println(f1(1), f2(2)) // 11 9
	fmt.Println(f1(3), f2(4)) // 12 8
	fmt.Println(f1(5), f2(6)) // 13 7

}

func calc(base int) (func(int) int, func(int) int) {
	//base变量在闭包内部共享的，是add、sub的共享变量
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}
