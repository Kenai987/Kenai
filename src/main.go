package main

import "fmt"

func main() {
	fmt.Println("开始了")
	//今天就开始慢慢计划

	var array [10]int

	var slice = array[5:6]

	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])
}
