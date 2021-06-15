package main

import "fmt"

func main() {
	p1 := Person{
		"杰森斯坦森",
		36,
	}
	fmt.Printf("%#v\n", p1)
}

// Person结构体
type Person struct {
	string
	int
}
