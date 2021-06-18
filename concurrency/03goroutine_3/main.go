package main

import "fmt"

func recv(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}

func main() {
	ch := make(chan int) //创建的是无缓冲的通道，必须要有goroutine去接收，否则会引发panic
	go recv(ch)          //启动goroutine接收值

	ch <- 10
	fmt.Println("发送成功", ch)
}
