package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func longestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}

	dp := make([]int, len(s)+1)
	dp[0] = 0
	dp[1] = 0

	ret := 0
	for i := 1; i < len(s); i++ {
		c := s[i]
		if c == ')' {
			if s[i-1] == '(' {
				dp[i+1] = dp[i-1] + 2
			} else {
				if i-dp[i]-1 >= 0 && s[i-dp[i]-1] == '(' {
					dp[i+1] = dp[i] + 2
					if i-dp[i]-1-1 >= 0 {
						dp[i+1] = dp[i+1] + dp[i-dp[i]-1]
					}
				}
			}
			ret = max(ret, dp[i+1])
		}
	}
	return ret
}

func main() {
	fmt.Println(longestValidParentheses(")()(())"))
}
