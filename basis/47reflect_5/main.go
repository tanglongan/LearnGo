package main

import (
	"fmt"
	"reflect"
)

//电脑结构体
type Computer struct {
	Brand    string `json:"name"`
	Model    string `json:"model"`
	Memory   int
	DiskType string
	DiskSize int
	Price    float32
}

//为结构体类型添加2个方法
func (c Computer) PlayGame() string {
	msg := "用电脑打游戏"
	fmt.Println(msg)
	return msg
}

func (c Computer) Program() string {
	msg := "用电脑来写代码"
	fmt.Println(msg)
	return msg
}

//打印类型的所有方法，同时通过反射机制调用方法
func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Printf("一共有%d个方法\n", t.NumMethod())

	//打印当前类型的所有方法
	for i := 0; i < t.NumMethod(); i++ {
		methodType := t.Method(i).Type
		fmt.Printf("Method Name: %s  Method: %s\n", t.Method(i).Name, methodType)
		args := []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func main() {
	c1 := Computer{
		Brand:    "华为",
		Model:    "Matebook 14",
		Memory:   16384,
		DiskType: "SSD",
		DiskSize: 524288,
		Price:    5678.9,
	}

	ct := reflect.TypeOf(c1)
	ck := ct.Kind()
	fmt.Println(c1)
	fmt.Printf("[Computer struct]==> Type：%v  Kind：%v\n=====================\n", ct, ck)

	//打印所有字段
	for i := 0; i < ct.NumField(); i++ {
		field := ct.Field(i)
		fmt.Printf("Name: %s Index: %d Type: %v JSON TAG: %v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}

	printMethod(c1)
}
