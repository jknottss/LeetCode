package main

import "fmt"

func sortArr(nums []int) []int {
	a := 0
	b := 1
	le := len(nums)
	res := make([]int, le)
	for i := 0; i < le; i++ {
		if nums[i]%2 == 0 {
			res[a] = nums[i]
			a += 2
		} else {
			res[b] = nums[i]
			b += 2
		}
	}
	return res
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(sortArr(arr))
}
