package main

import "fmt"

func main() {
	var price float64
	price = 23.23
	fmt.Println(price)

	var age = 23.3
	fmt.Println(age, price)

	var flag bool
	fmt.Println(flag)
	// 定义一个map 此时为nil
	var userMap map[string]int
	// 初始化
	userMap = make(map[string]int)

	userMap["a"] = 1
	userMap["b"] = 2

	fmt.Println(userMap)

	// ide都展示出来了
	const (
		a = iota
		b = iota
		c = 3
		d = iota
	)

	fmt.Println(a, b, c, d)
}
