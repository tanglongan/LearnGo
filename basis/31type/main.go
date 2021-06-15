package main

import "fmt"

type myInt int     //自定义类型，type后面跟的是类型
type yourInt = int //类型别名

func main() {
	var n myInt = 100
	fmt.Println(n)

	var c rune = '中' //rune本质上是int32的别名
	fmt.Println(c)
	fmt.Printf("%T\n", c)

}
