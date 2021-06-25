package main

import (
	"fmt"
	"math"
)

//定义“矩形”，并且实现ShapeDesc接口
type ShapeDesc interface {
	Area() float64
	Perimeter() float64
}

type rectangle struct {
	H, W float64
}

func (r rectangle) Area() float64 {
	return r.H * r.W
}

func (c rectangle) Perimeter() float64 {
	return 2 * (c.H + c.W)
}

//定义“圆”，并且实现ShapeDesc接口
type circle struct {
	R float64
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

//打印ShapeDesc类型的值
func printShape(s ShapeDesc) {
	switch s.(type) {
	case circle:
		fmt.Println("Shape is a circle")
	case rectangle:
		fmt.Println("Shape is a rectangle")
	}

	fmt.Println("Area: ", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func main() {
	var s1, s2 ShapeDesc
	s1 = rectangle{H: 2, W: 5}
	s2 = circle{R: 2}

	printShape(s1)
	printShape(s2)
}
