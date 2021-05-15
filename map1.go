package main

import "fmt"

func main() {
	// 创建 map
	scores := make(map[string]int)
	scores["Tom"] = 90
	scores["Jerry"] = 80

	// 遍历所有 map 的 keys
	for person := range scores {
		fmt.Println(person, "=", scores[person])
	}

	// 判断某个元素是否存在
	tomScore, ok := scores["Tom"]
	if ok {
		fmt.Println("Tom", "=", tomScore)
	}

	// 删除元素
	delete(scores, "Jerry")

	// 遍历所有 map 的 keys
	for person := range scores {
		fmt.Println(person, "=", scores[person])
	}
}
