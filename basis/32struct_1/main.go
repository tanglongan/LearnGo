package main

import "fmt"

//结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型的变量p
	var p person

	//为字段赋值
	p.name = "习近平"
	p.age = 60
	p.gender = "男"
	p.hobby = []string{"篮球", "足球"}
	fmt.Println(p)

}
