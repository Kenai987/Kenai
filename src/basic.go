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

	var userMap map[string]int
	userMap = make(map[string]int)

	userMap["a"] = 1
	userMap["b"] = 2

	fmt.Println(userMap)
}
