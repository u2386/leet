package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	p := m + n - 1
	i, j := m-1, n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[p] = nums1[i]
			i--
		} else {
			nums1[p] = nums2[j]
			j--
		}
		p--
	}

	for i = 0; i <= j; i++ {
		nums1[i] = nums2[i]
	}
}

func main() {
	a := []int{1, 2, 10, 0, 0, 0}
	b := []int{2, 5, 6}
	merge(a, 3, b, 3)
	fmt.Println(a)
}
