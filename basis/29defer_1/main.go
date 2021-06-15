package main

import "fmt"

//defer语句会将跟随的语句进行延迟处理。
//在defer归属的函数即将返回时，将延迟处理的语句按defer定义的逆序进行执行
//也就是说，先被defer的语句最后被执行，最后被defer的语句，最先被执行

/*
 * 第一步：返回值赋值
 * 第二步：defer
 * 第三步：真正的RET指令返回
 ********/
func main() {
	testDefer1()
	fmt.Println("=========================")

	fmt.Println(f1()) // 5
	fmt.Println(f2()) // 6
	fmt.Println(f3()) // 5
	fmt.Println(f4()) // 5
	fmt.Println(f5()) // 5
}

func testDefer1() {
	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++ //修改的是x，不是返回值
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 //返回x
}

func f3() (y int) {
	x := 5
	defer func() {
		x++ //修改的是x
	}()
	return x //返回值=y=x=5
}

func f4() (x int) {
	defer func(x int) {
		x++ //改变的是函数的副本
	}(x)

	return 5 //返回值=x=5
}

func f5() (x int) {
	defer func(x int) int {
		x++
		return x //这里的值并没有被接收
	}(x)

	return 5
}
