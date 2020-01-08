package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	dp := make([]int, len(nums)+1)
	dp[1] = nums[0]

	for i := 1; i < len(nums); i++ {
		dp[i+1] = max(dp[i], dp[i-1]+nums[i])
	}
	return dp[len(dp)-1]
}

func main() {
	arr := []int{2, 7, 9, 3, 1}
	fmt.Println(rob(arr))
}
