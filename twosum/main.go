package main

import "fmt"

func main() {
	i := []int{3, 3, 4, 15}
	t := 7

	fmt.Println(twoSum(i, t))

}

/*func twoSum(nums []int, target int) []int {*/
	/*//init map and empty array for return*/
	/*s := []int{}*/
	/*m := map[int]int{}*/

	/*for i := 0; i < len(nums); i++ {*/
		/*//set the value as the key and the index as the value in map m*/
		/*_, ok := m[nums[i]]*/
		/*//check the value exists, if it exists we need to check that the double isn't the result, if it is we can return early, if not we can ignore*/
		/*if ok {*/
			/*c := nums[i] * 2*/
			/*if c == target {*/
				/*s = []int{m[nums[i]], i}*/
				/*return s*/
			/*}*/
		/*}*/
		/*m[nums[i]] = i*/
	/*}*/
	/*//key and value of map m*/
	/*for k, v := range m {*/
		/*//if r + b = t; t - b = r; r is a key in our Map, if it exists and it is not at the same index then we know it's our answer*/
		/*r := target - k*/
		/*//check value exists*/
		/*_, ok := m[r]*/
		/*if ok {*/
			/*//make sure we are not returning the same value we are looking at*/
			/*if m[r] != v {*/
				/*s = []int{m[r], v}*/
			/*}*/
		/*}*/

	/*}*/
	/*return s*/

/*}*/

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currNum := range nums {
		if requiredIdx, isPresent := indexMap[target-currNum]; isPresent {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currNum] = currIndex
	}
	return []int{}
}
