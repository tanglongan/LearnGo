package main

import "fmt"

/*
 * 结构体是值类型
 **/
type person struct {
	name   string
	gender string
}

//Go语言中函数参数永远是拷贝
func f1(x person) {
	x.gender = "女" //修改的是副本的gender，因此不影响外面对象的属性值
}

func f2(x *person) {
	(*x).gender = "女" //根据内存地址找到原变量，修改的就是原来的变量，也可以是：x.gender = "女"
}

func main() {

	var p person
	p.name = "周星驰"
	p.gender = "男"
	f1(p)
	fmt.Println(p) //{周星驰 男}
	f2(&p)
	fmt.Println(p) //{周星驰 女}

	var p2 = new(person)
	p2.name = "特朗普"
	p2.gender = "男"
	fmt.Printf("%T\n", p2) //*main.person
	fmt.Printf("%p\n", p2) //p2保存的值就是一个内存地址
	fmt.Println(p2)        //&{ }

	//2：结构体指针
	//2.1、通过key-value初始化
	var p3 = &person{
		name:   "拜登",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)

	//2.2、使用值列表的初始化。值的顺序必须和结构体定义的字段顺序一致
	p4 := &person{
		"奥巴马", "男",
	}
	fmt.Printf("%#v\n", p4)

}
