package main

import "fmt"

/*
 * chan<- int：是一个只写单向的通道（只能对其写入int值），可以对其执行发送操作但是不能执行接收操作
 * <-chan int：是一个只读单向的通道（只能对其读取int值），可以对其执行接收操作但是不能执行发送操作
 */
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}
