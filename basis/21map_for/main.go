package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

//按照Key顺序进行遍历输出map
func main() {
	//初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	//数据初始化
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("stu%02d", i) //生成stu开头的字符串
		value := rand.Intn(100)          //生成0~99的随机数
		scoreMap[key] = value
	}

	//取出所有key存入切片keys中
	var keys = make([]string, 0, 200)
	for key := range scoreMap {
		keys = append(keys, key)
	}

	//对切片进行排序
	sort.Strings(keys)

	//按照排序后的key遍历map
	for _, key := range keys {
		fmt.Println(key, scoreMap[key])
	}

}
