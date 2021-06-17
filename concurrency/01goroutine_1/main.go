package main

import (
	"fmt"
	"sync"
)

//通过WaitGroup工具进行goroutine之间的同步
var wg sync.WaitGroup

func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1) //启动一个goroutine就加1
		go hello(i)
	}
	wg.Wait() //等待所有goroutine都结束执行
}

func hello(i int) {
	defer wg.Done() //goroutine停止就计数减1
	fmt.Println("Hello goroutine===>", i)
}
