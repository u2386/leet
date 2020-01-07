package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	dp := make([][]int, len(s))
	for i := range s {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	max := 1
	start := 0
	for j := 1; j < len(s); j++ {
		for i := 0; i < j; i++ {
			if s[i] == s[j] {
				if j-i < 3 {
					dp[i][j] = 1
				} else {
					dp[i][j] = dp[i+1][j-1]
				}

			} else {
				dp[i][j] = 0
			}

			if dp[i][j] == 1 {
				cur := j - i + 1
				if cur > max {
					start = i
					max = cur
				}
			}
		}
	}
	return s[start : start+max]
}

func main() {
	fmt.Println(longestPalindrome("babad"))
}
