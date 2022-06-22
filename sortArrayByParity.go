package main

import "fmt"

func sortArrayByParity(nums []int) []int {
	firstPointer := 0
	secondPointer := len(nums) - 1
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			res[firstPointer] = nums[i]
			firstPointer++
		} else {
			res[secondPointer] = nums[i]
			secondPointer--
		}
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sortArrayByParity(arr))
}
