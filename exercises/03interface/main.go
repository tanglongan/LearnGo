package main

import "fmt"

/*
 * interface：关键字
 * interface{}：空接口
 */
func main() {

	m1 := make(map[string]interface{}, 16)
	m1["name"] = "TANG"
	m1["age"] = 30
	m1["merried"] = false
	m1["hobby"] = [...]string{"打羽毛球", "中国象棋", "刷B站"}
	fmt.Println(m1)

	show(nil)
	show(m1)
}

//空接口作为函数参数
func show(x interface{}) {
	fmt.Printf("type:%T value:%v\n", x, x)
}
