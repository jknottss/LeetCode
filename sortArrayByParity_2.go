package main

import "fmt"

func sort(nums []int) []int {
	idx := 0
	for i := range nums {
		if nums[i]%2 == 0 {
			nums[i], nums[idx] = nums[idx], nums[i]
			idx++
			fmt.Println(nums)
		}
	}
	return nums
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sort(arr))
}
