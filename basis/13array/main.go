package main

import "fmt"

/*
 * 数组是同一种数据类型的元素的集合。
 * 在Go语言中，数据中从声明就确定，使用时可以修改数据成员，但是数据大小不可变化。
 * 定义格式：var 数据变量名 [元素数量] T
 */
func main() {
	testArray1()
	fmt.Println("==========================")
	testArray2()
	fmt.Println("==========================")
	testArray3()
	fmt.Println("==========================")
	testPrint()
	fmt.Println("==========================")
	testTwoArray()
	fmt.Println("==========================")
	testModifyArray()
}

//方式一：初始化数组时可以使用初始化列表来设置数组元素的值
func testArray1() {
	var testArray [3]int
	var numArray = [3]int{1, 2}
	var cityArray = [3]string{"北京", "上海", "深圳"}

	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
}

//方式二：按照上面的方法每次都要确保提供的初始值和数组长度一致，一般情况下，我们可以让编译器根据初始值的个数自行推断数组的长度
func testArray2() {
	var testArray [3]int
	var numArray = [...]int{1, 2}
	var cityArray = [...]string{"北京", "上海", "深圳", "广州"}

	fmt.Println(testArray)
	fmt.Println(numArray)
	fmt.Println(cityArray)
}

//方式三：使用指定索引值的方式来初始化数组
func testArray3() {
	a := [...]int{1: 1, 3: 3}
	fmt.Println(a)
	fmt.Printf("type of a:  %T\n", a)
}

//遍历数组
func testPrint() {
	var a = [...]string{"北京", "上海", "深圳", "广州"}
	fmt.Println("-------遍历方式1-----")
	//方式一：
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
	fmt.Println("-------遍历方式2-----")
	//方式二：
	for _, value := range a {
		fmt.Println(value)
	}
}

//二维数组的声明和遍历
func testTwoArray() {
	a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"天津", "西安"},
	}
	fmt.Println(a)
	fmt.Println(a[2][1])

	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}

/*
 * 下面3个函数用来说明：
 *  数组是值类型，复制和传参会复制整个数组。因此改变副本的值，不会改变自身的值。
 */
func modifyArray1(x [3]int) {
	x[0] = 100
}

func modifyArray2(x [3][2]int) {
	x[2][0] = 100
}

func testModifyArray() {
	a := [3]int{10, 20, 30}
	modifyArray1(a) //在modifyArray1函数中修改的是a的副本
	fmt.Println(a)

	b := [3][2]int{
		{1, 1},
		{2, 2},
		{3, 3},
	}

	modifyArray2(b)
	fmt.Println(b)
}
