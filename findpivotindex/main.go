package main

import "fmt"

func pivotIndex(nums []int) int {

	left := 0
	right := 0

	for i, v := range nums {
		if i != 0 {
			right += v
		}
	}

	for i, v := range nums {
		if left == right {
			return i
		}
		if i < len(nums)-1 {
			left += v
			right -= nums[i+1]
		}
	}
	return -1
}

func main() {
	nums := []int{1, 7, 3, 6, 5, 6}
	fmt.Println(pivotIndex(nums))
}
