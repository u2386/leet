package main

import "fmt"

func isPalindrome(s string) bool {
	if len(s) <= 1 {
		return true
	}

	i, j := 0, len(s)-1
	var ci, cj byte

	for i < len(s) && j >= 0 && i < j {
		for i < len(s) {
			ci = s[i]
			if ci >= 'a' && ci <= 'z' || ci >= '0' && ci <= '9' {
				break
			} else if ci >= 'A' && ci <= 'Z' {
				ci = ci + 32
				break
			}
			i++
		}

		for j >= 0 {
			cj = s[j]
			if cj >= 'a' && cj <= 'z' || cj >= '0' && cj <= '9' {
				break
			} else if cj >= 'A' && cj <= 'Z' {
				cj = cj + 32
				break
			}
			j--
		}

		if !(i <= j) {
			return true
		}
		if ci != cj {
			return false
		}
		i++
		j--
	}

	return true
}

func main() {
	s := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(s))
}
