package main

import "fmt"

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxProduct(nums []int) int {
	imax := 1
	imin := 1

	a := uint(1) << 63
	ret := int(a)
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			imax, imin = imin, imax
		}
		imax = max(imax*nums[i], nums[i])
		imin = min(imin*nums[i], nums[i])
		if imax > ret {
			ret = imax
		}
	}
	return ret
}

func main() {
	arr := []int{2, 3, -2, 4}
	fmt.Println(maxProduct(arr))

}
