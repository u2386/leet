package main

import "fmt"

func backtrace(s string, start int, dp [][]int, path []string, ret [][]string) [][]string {
	if start == len(s) {
		ret = append(ret, path)
		return ret
	}

	for i := start; i < len(s); i++ {
		if dp[start][i] != 1 {
			continue
		}

		path = append(path, s[start:i+1])
		ret = backtrace(s, i+1, dp, append(path[:0:0], path...), ret)
		path = path[:len(path)-1]
	}
	return ret
}

func partition(s string) [][]string {
	ret := [][]string{}
	if len(s) == 0 {
		return ret
	}

	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}

	for j := 0; j < len(s); j++ {
		for i := 0; i <= j; i++ {
			if s[i] == s[j] && (j-i <= 2 || dp[i+1][j-1] == 1) {
				dp[i][j] = 1
			}
		}
	}
	return backtrace(s, 0, dp, []string{}, ret)
}

func main() {
	s := "aab"
	fmt.Println(partition(s))
}
