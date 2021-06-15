package main

import "fmt"

type Sayer interface {
	say()
}

type dog struct {
}

type cat struct {
}

func (d dog) say() {
	fmt.Println("汪汪汪")
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func main() {
	d := &dog{}
	d.say()

	c := &cat{}
	c.say()
}
