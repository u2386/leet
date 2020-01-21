package main

import (
	"fmt"
)

func min(a, b, c int) int {
	tmp := a
	if b < a {
		tmp = b
	}

	if tmp < c {
		return tmp
	}
	return c
}

func minDistance(word1 string, word2 string) int {

	var dp func(i, j int) int
	dp = func(i, j int) int {
		if i == -1 {
			return j + 1
		} else if j == -1 {
			return i + 1
		}

		if word1[i] == word2[j] {
			return dp(i-1, j-1)
		}

		return min(
			dp(i, j-1)+1,
			dp(i-1, j)+1,
			dp(i-1, j-1)+1,
		)

	}

	return dp(len(word1)-1, len(word2)-1)
}

func main() {
	fmt.Println(minDistance("horse", "ros"))
}
