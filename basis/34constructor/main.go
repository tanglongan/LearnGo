package main

import "fmt"

func main() {
	p1 := newPerson("莱昂纳德", 30)
	fmt.Println(p1)
}

type person struct {
	name string
	age  int
}

//自定义一个构造方法，用于创建person结构体实例
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
