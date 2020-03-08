package main

import (
	"fmt"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func coinChange(coins []int, amount int) int {
	if len(coins) == 0 || amount <= 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0

	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

func main() {
	coins := []int{186, 419, 83, 408}
	amount := 6249
	fmt.Println(coinChange(coins, amount))
}
