package main

import "fmt"

func main() {
	nums := []int{1, 3, 4, 2, 2}
	fmt.Println(findDuplicate(nums))
}

func findDuplicate(nums []int) int {
	defer printDefer(nums)
	for index := range nums {
		fmt.Println(index)
	}

	return 1
}

func printDefer(nums []int) {
	fmt.Println(nums)
}
