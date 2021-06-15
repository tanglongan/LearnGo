package main

import "fmt"

/*
 * 1.可变参数函数：可变参数需要放在函数参数最后一个参数上
 * 2.
 */
func main() {

	fmt.Println(testIntSum(10, 1, 2, 3))
	fmt.Println(testIntSum(10, 1, 2, 3, 5, 6, 7))

}

/*
 * 可变参数
 */
func testIntSum(x int, y ...int) int {
	sum := x
	for _, v := range y {
		sum = sum + v
	}
	return sum
}
