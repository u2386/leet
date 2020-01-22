package main

import "fmt"

func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	a, b := 1, 1
	for i := 1; i < len(s); i++ {
		if s[i] != '0' {
			if s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26" {
				a, b = a+b, a
			} else {
				b = a
			}
			continue
		}

		if !(s[i-1:i+1] >= "10" && s[i-1:i+1] <= "26") {
			return 0
		}
		a = b
	}
	return a
}

func main() {
	fmt.Println(numDecodings("12021"))
}
