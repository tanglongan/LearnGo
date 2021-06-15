package main

import "fmt"

// append()为切片追加元素

func main() {

	s1 := []string{"北京", "上海", "深圳"}

	s1[2] = "广州" //要注意索引，否则会引起索引越界
	fmt.Println(s1)

	//调用append函数在切片末尾添加元素
	s1 = append(s1, "西安")
	fmt.Println(s1)

	s2 := append(s1, "武汉")
	fmt.Println(s2)

	ss := []string{"杭州", "苏州", "长沙"}
	s2 = append(s2, ss...) // ...表示将切片元素拆分开
	fmt.Println(s2)

}
