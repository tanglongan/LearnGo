package main

import (
	"fmt"
	"reflect"
)

/*
 * 先判断值类型，然后取值
 **/
func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func main() {
	var a float64 = 3.1415
	var b int64 = 100
	c := reflect.ValueOf(10)
	reflectValue(a)               //type is float64, value is 3.141500
	reflectValue(b)               //type is int64, value is 100
	fmt.Printf("type c :%T\n", c) //type c :reflect.Value
}
