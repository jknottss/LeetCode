package main

func maxSubArray(nums []int) int {
	sum := 0
	res := math.MinInt32
	for _, n := range nums {
		sum = n + max(sum, 0)
		res = max(sum, res)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
