package main

import "fmt"

func minWindow(s string, t string) string {
	ret := ""
	var need, watch [256]int

	for _, c := range t {
		need[c-'A'] = 1
		watch[c-'A']++
	}

	minW := len(s)
	match := len(t)

	left, right := 0, 0
	for right < len(s) {
		if need[s[right]-'A'] == 1 {
			if watch[s[right]-'A'] > 0 {
				match--
			}
			watch[s[right]-'A']--
		}

		for match == 0 {
			if right-left < minW {
				minW = right - left
				ret = s[left:right+1]
			}

			if need[s[left]-'A'] == 1 {
				watch[s[left]-'A']++
				if watch[s[left]-'A'] > 0 {
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
	s := "ADOBECODEBANC"
	t := "ABC"

	fmt.Println(minWindow(s, t))
}
