package main

import (
	"fmt"
	"reflect"
)

//自定义类型
type MyString string

//自定义结构体
type Student struct {
	Name      string
	Age       int8
	Interests []string
}

func main() {
	s1 := Student{
		Name:      "韦小宝",
		Age:       30,
		Interests: []string{"撩妹子", "吹牛逼", "扮太监"},
	}

	st1 := reflect.TypeOf(s1)
	sk1 := st1.Kind()
	fmt.Println("Type: ", st1)                        //Type:  main.Student
	fmt.Println("Kind: ", sk1, sk1 == reflect.Struct) //Kind:  struct true

	fmt.Println("====================")
	s2 := "Hello World"
	st2 := reflect.TypeOf(s2)
	sk2 := st2.Kind()
	fmt.Println("Type: ", st2)                        //Type:  string
	fmt.Println("Kind: ", sk2, sk2 == reflect.String) //Kind:  string true

	fmt.Println("====================")
	s3 := &Student{
		Name:      "张无忌",
		Age:       28,
		Interests: []string{"乾坤大挪移", "吸星大法"},
	}
	st3 := reflect.TypeOf(s3)
	sk3 := st3.Kind()
	fmt.Println("Type: ", st3)                     //Type:  *main.Student
	fmt.Println("Kind: ", sk3, sk3 == reflect.Ptr) //Kind:  ptr true

}
