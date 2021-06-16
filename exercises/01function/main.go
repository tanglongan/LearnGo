package main

import (
	"bytes"
	"fmt"
)

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9} // 等价于arr := [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T\n", arr)                    // [9]int
	sum(arr[:]...)

	a := []string{"a", "b", "c"}
	b := []string{"d", "e", "f"}

	//可变参数变量是一个包含所有参数的切片，如果要将这个含有可变参数的变量传递给下一个可变参数函数，可以在传递时给可变参数变量后面添加...，这样就可以将切片中的元素进行传递，而不是传递可变参数变量本身。
	//将切片的元素打散，一个个地传入函数
	a = append(a, b...)
	fmt.Println(a)

	fmt.Println(printTypeValue(100, "str", true))
}

//定义可变参数的函数，必须是函数最后一个参数
func sum(nums ...int) {
	for _, v := range nums {
		fmt.Println(v)
	}
}

//遍历参数，并将所有参数信息拼接成字符串返回
func printTypeValue(slist ...interface{}) string {
	var b bytes.Buffer
	for _, s := range slist {
		str := fmt.Sprintf("%v", s)
		var typeString string
		switch s.(type) {
		case bool:
			typeString = "bool"
		case string:
			typeString = "string"
		case int:
			typeString = "int"
		}

		b.WriteString("value: ")
		b.WriteString(str)
		b.WriteString(" type: ")
		b.WriteString(typeString)
		b.WriteString("\n")
	}
	return b.String()
}
