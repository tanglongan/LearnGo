package main

import "fmt"

/*
 * 方法接收者
 */
func main() {
	p1 := NewPerson("库里", 33)
	p1.Dream()

	p1.SetAge(32)
	fmt.Println(p1)
	p1.SetAge2(30)
	fmt.Println(p1)
}

//Person结构体
type Person struct {
	name string
	age  int8
}

//Person结构体的构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//为Person定义一个方法： Dream()
func (p Person) Dream() {
	fmt.Printf("%s的梦想就是学好Go语言\n", p.name)
}

//设置p的年龄，指针接收者
func (p *Person) SetAge(newAge int8) {
	p.age = newAge
}

//设置p的年龄，值接收者
func (p Person) SetAge2(newAge int8) {
	//p.age = newAge
}
