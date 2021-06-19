package main

import "fmt"

/*
 * 同一个结构体可以实现多个接口
 * 接口还可以嵌套组合
 *
 *
 */
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

//cat实现了mover接口
func (c *cat) move() {
	fmt.Println("走猫步.....")
}

//cat实现了eater接口
func (c *cat) eat(food string) {
	fmt.Printf("猫吃%s", food)
}

func main() {
	c := cat{name: "蓝猫", feet: 4}
	c.move()
	c.eat("小黄鱼")

}
