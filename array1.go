package main

import "fmt"

func main() {
	// 普通的方式定义数组
	var people [3]string
	people[0] = "hello"
	people[1] = "world"
	people[2] = "haha"

	for i := 0; i < len(people); i++ {
		fmt.Printf("%s\n", people[i])
	}

	// 定义数组的同时赋值
	numbers := [...]int{2, 4, 6, 8, 10}
	for j := 0; j < len(numbers); j++ {
		fmt.Printf("%d\n", numbers[j])
	}
}
