package main

import "fmt"

/*
 * 取地址操作符&和取值操作符*是一对互补的操作：&取出地址；*根据地址取出地址指向的值
 * 对变量进行取地址&操作，可以获取到这个变量的指针变量
 * 指针变量的值是指针地址
 * 对指针变量进行取值*操作，可以获取指针变量指向的原变量的值
 */
func main() {

	testPointer1()
	testPointer2()

}

func testPointer1() {
	fmt.Println("=========================")
	n := 18
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p)
}

func testPointer2() {
	fmt.Println("=========================")

	//var a *int
	//此处有问题，此时变量并没有分配内存空间，因此直接赋值是错误的
	//*a = 100
	//fmt.Println(a)

	//new函数申请一个内存地址
	var b = new(int)
	*b = 200
	fmt.Println(*b, &b)
}
