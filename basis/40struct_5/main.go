package main

import "fmt"

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{
			name: "二哈",
		},
	}

	d1.move()
	d1.wang()
}

type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}
