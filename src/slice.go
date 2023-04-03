package main

import "fmt"

func main() {
	var array [10]int

	var slice = array[5:7]

	fmt.Println("lenth of slice: ", len(slice))
	fmt.Println("capacity of slice: ", cap(slice))
	fmt.Println(&slice[0] == &array[5])

	var slice2 []int
	slice2 = append(slice2, 1, 2, 3)

	newSlice := AddElement(slice2, 4)
	fmt.Println(&slice2[0] == &newSlice[0])

}

func AddElement(slice []int, e int) []int {
	return append(slice, e)
}
