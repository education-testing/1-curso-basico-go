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

	testContinue(nums)
	testBreak(nums)

	return 1
}

func printDefer(nums []int) {
	fmt.Println(nums)
}

func testContinue(nums []int) {
	fmt.Println("Starting continue test")
	for i := 0; i < len(nums); i++ {

		if nums[i] == 2 {
			continue
		}

		fmt.Println(nums[i])
	}
}

func testBreak(nums []int) {
	fmt.Println("Starting break test")
	for i := 0; i < len(nums); i++ {
		if nums[i] == 2 {
			break
		}

		fmt.Println(nums[i])
	}
}
