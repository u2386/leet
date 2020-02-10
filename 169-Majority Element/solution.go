package main

import "fmt"

func majorityElement(nums []int) int {
	count := 1
	n := nums[0]

	for i := 1; i < len(nums); i++ {
		if count == 0 {
			n = nums[i]
			count = 1
			continue
		}

		if n == nums[i] {
			count++
		} else {
			count--
		}
	}
	return n
}

func main() {
	arr := []int{7, 7, 5, 7, 5, 1, 5, 7, 5, 5, 7, 7, 7, 7, 7, 7}
	fmt.Println(majorityElement(arr))
}
