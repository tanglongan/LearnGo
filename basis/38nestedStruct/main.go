package main

import "fmt"

// Address地址结构体
type Address struct {
	Province string
	City     string
}

// User用户结构体
type User struct {
	Name   string
	Gender string
	Address
}

func main() {

	user1 := User{
		Name:   "马化腾",
		Gender: "男",
		Address: Address{
			Province: "广东省",
			City:     "深圳市",
		},
	}

	fmt.Printf("user1=%#v\n", user1)
}
