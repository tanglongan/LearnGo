package main

import "fmt"

// if条件判断
func main() {

	age := 30

	if age > 60 {
		fmt.Println("老年人")
	} else if age > 40 {
		fmt.Println("中年人")
	} else if age > 18 {
		fmt.Println("青年人")
	} else {
		fmt.Println("青少年")
	}

}
