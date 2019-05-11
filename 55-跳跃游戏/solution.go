package main

import "fmt"

func canJump(nums []int) bool {
	length := len(nums)
	max := 0
	for i, n := range(nums) {
		if i > max {
			return false
		} else {
			if i + n > max {
				max = i + n
			}
			if max >= length - 1 {
				return true
			}
		}
	}

	return false
}


func main() {
	n := []int{2,3,1,1,4}
	fmt.Println(canJump(n))
}
