package main

import (
	"encoding/json"
	"fmt"
)

// Student学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

// Class班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	//准备数据
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}

	for i := 0; i < 3; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}

	//JSON序列化：结构体数据 -->JSON
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Json marshal failed")
		return
	}

	fmt.Printf("json: %s\n", data)

	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"}]}`
	c1 := &Class{}
	er := json.Unmarshal([]byte(str), c1)
	if er != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
