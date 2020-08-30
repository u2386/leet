package main

import "fmt"

func isMatch(s string, p string) bool {
	dp := make([][]bool, len(p)+1)
	for i := 0; i < len(p)+1; i++ {
		dp[i] = make([]bool, len(s)+1)
	}
	dp[0][0] = true

	for i := 1; i <= len(p); i++ {
		for j := 0; j <= len(s); j++ {
			if p[i-1] == '*' {
				dp[i][j] = dp[i-2][j]
				if j > 0 && (p[i-2] == '.' || p[i-2] == s[j-1]) {
					dp[i][j] = dp[i][j] || dp[i][j-1]
				}
			} else if j > 0 && (p[i-1] == '.' || p[i-1] == s[j-1]) {
				dp[i][j] = dp[i-1][j-1]
			}
		}
	}
	return dp[len(dp)-1][len(dp[0])-1]
}

func main() {
	fmt.Println(isMatch("aa", "a*"))
}
