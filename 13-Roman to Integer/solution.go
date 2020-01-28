package main

import (
	"fmt"
)

func romanToInt(s string) int {
	R := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	num := []int{}
	i := 0
	for i < len(s)-1 {
		v0, v1 := R[s[i]], R[s[i+1]]
		if v0 < v1 {
			num = append(num, v1-v0)
			i += 2
			continue
		} else {
			num = append(num, v0)
			i++
		}
	}
	if i < len(s) {
		num = append(num, R[s[i]])
		i++
	}

	ret := 0
	for _, n := range num {
		ret += n
	}
	return ret
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
}
