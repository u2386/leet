package main

import (
	"fmt"
)

func numDistinct(s string, t string) int {
	dp := make([][]int, len(t)+1)
	for i := 0; i < len(t)+1; i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for i := 0; i < len(s)+1; i++ {
		dp[0][i] = 1
	}

	for i := 0; i < len(t); i++ {
		for j := 0; j < len(s); j++ {
			if t[i] != s[j] {
				dp[i+1][j+1] = dp[i+1][j]
			} else {
				dp[i+1][j+1] = dp[i][j] + dp[i+1][j]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	S := "rabbbit"
	T := "rabbit"
	fmt.Println(numDistinct(S, T))
}
