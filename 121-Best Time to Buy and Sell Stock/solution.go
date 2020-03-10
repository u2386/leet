package main

import (
	"fmt"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	dp := make([][]int, len(prices))
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, 2)
		if i == 0 {
			dp[i][0] = 0
			dp[i][1] = -prices[i]
			continue
		}

		dp[i][0] = max(dp[i-1][0], dp[i-1][1]+prices[i])
		dp[i][1] = max(dp[i-1][1], -prices[i])
	}
	fmt.Println(dp)
	return dp[len(dp)-1][0]
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(arr))
}
