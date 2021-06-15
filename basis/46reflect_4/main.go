package main

import (
	"fmt"
	"reflect"
)

/*
 * 1.想要在函数中通过反射修改变量的值。需要注意函数传递的是值拷贝，必须传递变量地址才能够修改变量的值，否则会引发panic
 * 2.反射中使用Elem()方法来获取指针对应的值
 */
func main() {
	var a int64 = 100
	reflectSetValue2(&a)
	fmt.Println(a)

	reflectSetValue1(a) //panic: 引发panic，必须通过变量地址才能修改变量的值
}

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(300)
	}
}
