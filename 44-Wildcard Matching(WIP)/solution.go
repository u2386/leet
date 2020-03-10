package main

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}

	dp := make([][]int, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		dp[i] = make([]int, len(s)+1)
		for j := 0; j < len(s)+1; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = 1
			} else if i == 0 {
				dp[i][j] = 0
			} else if j == 0 {
				if p[i-1] == '?' && dp[i-1][j] == 1 {
					dp[i][j] = 1
				} else {
					dp[i][j] = 0
				}
			} else {
				dp[i][j] = dp[i-1][j]
				if i == j && p[i-1] != s[j-1] {
					return false
				}
			}
		}
		fmt.Println(dp)
	}

	return dp[len(p)][len(s)] == 1
}

func main() {
	s := "adceb"
	p := "adc?b"
	fmt.Println(isMatch(s, p))
}
