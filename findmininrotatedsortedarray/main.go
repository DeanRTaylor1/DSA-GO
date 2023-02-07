package main

import (
	"fmt"
	"math"
)

//Suppose an array of length n sorted in ascending order is rotated between 1 and n times. For example, the array nums = [0,1,2,4,5,6,7] might become:

//[4,5,6,7,0,1,2] if it was rotated 4 times.
//[0,1,2,4,5,6,7] if it was rotated 7 times.
//Notice that rotating an array [a[0], a[1], a[2], ..., a[n-1]] 1 time results in the array [a[n-1], a[0], a[1], a[2], ..., a[n-2]].

//Given the sorted rotated array nums of unique elements, return the minimum element of this array.

//You must write an algorithm that runs in O(log n) time.

func main() {
	nums := []int{11, 13, 15, 17}

	fmt.Println(findMin(nums))
}

func findMin(nums []int) int {
	res := nums[0]
	left := 0

	right := len(nums) - 1

	for left < right {
		// if the pointers overlap we need to exit the loop

		fmt.Println(nums[left], nums[right], res)
		if nums[left] < nums[right] {
			res = int(math.Min(float64(res), float64(nums[left])))
      break
		}
		m := (left + right) / 2
		fmt.Println(res, m, nums[m])
		res = int(math.Min(float64(res), float64(nums[m])))
		fmt.Println(res, m, nums[m])

		if nums[m] >= nums[left] {
			left = m + 1; 
    } else {
			right = m 
		}
	}

	return res
}
