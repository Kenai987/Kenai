package main

import "fmt"

func main() {
	a := 5
	b := 6

	fmt.Println(a, b)
	a, b = change(a, b)
	fmt.Println(a, b)

	max := max(a, b)
	fmt.Println(max)

	changeFunc(1, 2, 4, 5, 6, 7, 8, 9, 9)

	point := 23
	point = pointFunc(&point)
	point = pointFunc(&point)
	point = pointFunc(&point)
	fmt.Println(point)

	deferFunc(1)

}

// 交换函数
func change(a int, b int) (newA int, newB int) {
	return b, a
}

/**
最大值函数
*/
func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// 可变参数
func changeFunc(arg ...int) int {
	for n := range arg {
		fmt.Println(n)
	}
	return arg[0]
}

// 指针函数
// 行参会拷贝实参，船指针会提高性能，但都是在原始对象上修改。
func pointFunc(num *int) int {
	*num = *num + 1
	return *num
}

// 延迟函数

func deferFunc(a int) int {
	defer fmt.Println(a)
	a++
	defer fmt.Println(a)
	a += 10
	defer fmt.Println(a)
	if a > 100000 {
		return a
	}
	defer deferFunc(a)
	return a
}
