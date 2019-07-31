package main

import "fmt"

func minWindow(s string, t string) string {
	ret := ""
	needs := make(map[byte]int)

	for _, c := range t {
		needs[byte(c)]++
	}

	minLength := len(s)
	match := len(t)

	left, right := 0, 0
	for right < len(s) {
		if _, ok := needs[s[right]]; ok {
			if needs[s[right]] > 0 {
				match--
			}
			needs[s[right]]--
		}

		for match == 0 {
			if right-left < minLength {
				minLength = right - left
				ret = s[left:right+1]
			}

			if _, ok := needs[s[left]]; ok {
				needs[s[left]]++
				if needs[s[left]] > 0 {
					match++
				}
			}
			left++
		}

		right++
	}

	return ret
}

func main() {
	s := "bba"
	t := "ab"

	fmt.Println(minWindow(s, t))
}
