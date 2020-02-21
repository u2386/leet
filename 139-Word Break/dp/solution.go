package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]struct{})
	for _, word := range wordDict {
		wordSet[word] = struct{}{}
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i < len(s)+1; i++ {
		for j := 0; j < i; j++ {
			if _, ok := wordSet[s[j:i]]; ok && dp[j] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(dp)-1]
}

func main() {
	s := "leetcode"
	wd := []string{"leet", "code"}
	fmt.Println(wordBreak(s, wd))
}
