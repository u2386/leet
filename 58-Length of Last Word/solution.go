package main

import (
	"fmt"
)

func lengthOfLastWord(s string) int {
	i := len(s) - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
	}

	j := i
	for ; j >= 0; j-- {
		if s[j] == ' ' {
			break
		}
	}
	return i - j
}

func main() {
	s := "a "
	fmt.Println(lengthOfLastWord(s))
}
