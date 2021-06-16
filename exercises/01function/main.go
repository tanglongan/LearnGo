package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T\n", arr)
	sum(arr[:]...)
}

func sum(nums ...int) {
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
}
