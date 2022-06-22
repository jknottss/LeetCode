package main

func maxSubArray(nums []int) int {
	sum := 0
	max := nums[0]
	for i := 0; i < len(nums); i++ {
		if sum >= 0 {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > max {
			max = sum
		}
	}
	return max
}
