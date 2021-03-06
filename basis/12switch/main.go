package main

import "fmt"

func main() {
	testSwitch1()
	testSwitch2()
	testSwitch3()
	testSwitch4()
}

//switch语句可以方便地对大量的值进行条件判断
func testSwitch1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入")
	}
}

//一个分支可以有多个值，多个case值之间使用英文逗号分隔
func testSwitch2() {
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 0:
		fmt.Println("偶数")
	}
}

//分支还可以使用表达式，这时候switch语句后面不需要再跟判断变量
func testSwitch3() {
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好~")
	}
}

//fallthrough语法可以执行满足条件的case的下一个case，是为了兼容C语言中的case设计
func testSwitch4() {
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	}
}
