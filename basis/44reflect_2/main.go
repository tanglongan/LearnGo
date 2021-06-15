package main

import (
	"fmt"
	"reflect"
)

//自定义类型
type MyInt int64

//自定义结构体person
type person struct {
	name string
	age  int
}

//自定义结构体book
type book struct {
	title string
}

//通过反射机制获取：类型、值
func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v kind: %v\n", t.Name(), t.Kind()) //这里Name()方法，如果是内置类型就返回其名称，否则返回空
}

func main() {
	var a *float32 //指针类型
	var b MyInt    //自定义类型
	var c rune     //类型别名
	var d = person{
		name: "迪丽热巴",
		age:  30,
	}
	var e = book{title: "Go语言高级教程"}

	reflectType(a) //type: kind:ptr
	reflectType(b) //type:myInt kind:int64
	reflectType(c) //type:int32 kind:int32
	reflectType(d) //type: person kind: struct
	reflectType(e) //type: book kind: struct

}
