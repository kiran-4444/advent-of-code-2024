package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(remove(slice, 0))
	fmt.Println(remove(slice, 1))
}

func remove(slice []int, index int) []int {
	sliceCopy := make([]int, len(slice))
	copy(sliceCopy, slice)
	return append(sliceCopy[:index], sliceCopy[index+1:]...)
}
