package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "lnj", "请输入人的姓名")
	age := flag.Int("age", 33, "请输入人的年龄")
	// 2.注册解析命令行参数
	flag.Parse()
	// 3.使用命令行参数
	// 注意flag对应方法返回的都是指针类型, 所以使用时必须通过指针访问
	fmt.Println("name = ", *name)
	fmt.Println("age = ", *age)
}
