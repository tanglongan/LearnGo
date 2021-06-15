package main

import "fmt"

/*
 * map是一种无序的key-value的数据结构，Go语言的map是引用类型
 * 定义格式： var varName map[kType]vType
 * 通过make函数分配内存
 */
func main() {

	// 定义变量，同时使用make函数初始化分配内存
	var m1 = make(map[string]int, 10)
	m1["中国"] = 100
	m1["美国"] = 90
	m1["英国"] = 70
	fmt.Println(m1)

	//根据Key获取值
	fmt.Println("============================")
	fmt.Println(m1["中国"])

	//如果不存在这个Key，返回对应值类型的零值
	fmt.Println("============================")
	fmt.Println(m1["法国"])

	//判断Key是否存在
	fmt.Println("============================")
	value, ok := m1["俄罗斯"]
	if !ok {
		fmt.Println("查无此Key")
	} else {
		fmt.Println(value)
	}

	//range遍历
	fmt.Println("============================")
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	//删除
	fmt.Println("============================")
	delete(m1, "美国")
	fmt.Println(m1)

}
