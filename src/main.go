package main

import (
	"fmt"
	"runtime"
)

func main() {
	//fmt.Println("开始了")
	////今天就开始慢慢计划
	//
	//var array [10]int
	//
	//var slice = array[5:6]
	//
	//fmt.Println("lenth of slice: ", len(slice))
	//fmt.Println("capacity of slice: ", cap(slice))
	//fmt.Println(&slice[0] == &array[5])

	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
