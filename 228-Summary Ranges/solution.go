package main

import (
	"fmt"
	"strconv"
	"strings"
)

func summaryRanges(nums []int) []string {
	ret := []string{}
	if len(nums) == 0 {
		return ret
	}

	i := 0
	for i < len(nums) {
		var sb strings.Builder
		sb.WriteString(strconv.Itoa(nums[i]))

		flag := false
		for {
			if i+1 < len(nums) && nums[i+1]-nums[i] == 1 {
				i++
				flag = true
				continue
			}
			if flag {
				sb.WriteString("->")
				sb.WriteString(strconv.Itoa(nums[i]))
			}
			break
		}
		ret = append(ret, sb.String())
		i++
	}
	return ret
}

func main() {
	arr := []int{0, 1, 2, 4, 5, 7}
	fmt.Println(summaryRanges(arr))
}
