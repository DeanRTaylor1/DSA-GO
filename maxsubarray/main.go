package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	fmt.Println(maxSubArray(nums))

}

func maxSubArray(nums []int) int {
	cSum := 0
	res := nums[0]

	for _, v := range nums {
		if cSum < 0 {
			cSum = 0
		}
		cSum += v
		res = int(math.Max(float64(res), float64(cSum)))
	}
	return res
}
