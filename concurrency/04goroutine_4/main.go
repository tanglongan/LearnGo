package main

import "fmt"

func main() {

	ch1 := make(chan int)
	ch2 := make(chan int)

	//开启goroutine将0~100发送到ch1通道中
	go func() {
		for i := 0; i < 10; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	//开启goroutine从ch1通道获取数据，并将数据的平方值放入ch2通道中
	go func() {
		for {
			i, ok := <-ch1
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()

	//从main goroutine中获取数据并打印
	for v := range ch2 {
		fmt.Println(v)
	}
}
