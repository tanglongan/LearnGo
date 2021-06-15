package main

import (
	"fmt"
	"strings"
)

/*
 * 统计一个字符串中每个单词出现的次数
 */
func main() {

	var testString = "How do you do "
	result := testCountSentence(testString)

	for stringworld, count := range result {
		fmt.Printf("%v has [%v]: %v\n", testString, stringworld, count)
	}

}

func testCountSentence(testString string) map[string]int {
	stringAfterop := strings.Split(testString, " ")
	mapDemo := make(map[string]int, 10)

	for _, v := range stringAfterop {
		if v == "" {
			continue
		}
		mapDemo[v]++
	}
	return mapDemo
}
