package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test01(t *testing.T) {
	intervals := [][]int{
		[]int{1, 3},
		[]int{6, 9},
	}
	newInterval := []int{2, 5}

	expected := [][]int{
		[]int{1, 5},
		[]int{6, 9},
	}
	assert.ElementsMatch(t, expected, insert(intervals, newInterval))
}

func Test02(t *testing.T) {
	intervals := [][]int{
		[]int{1, 2},
		[]int{3, 5},
		[]int{6, 7},
		[]int{8, 10},
		[]int{12, 16},
	}
	newInterval := []int{4, 8}
	expected := [][]int{
		[]int{1, 2},
		[]int{3, 10},
		[]int{12, 16},
	}
	assert.ElementsMatch(t, expected, insert(intervals, newInterval))
}

func Test03(t *testing.T) {
	intervals := [][]int{
		[]int{2, 4},
		[]int{5, 7},
		[]int{8, 10},
		[]int{11, 13},
	}
	newInterval := []int{3, 8}
	expected := [][]int{
		[]int{2, 10},
		[]int{11, 13},
	}
	assert.ElementsMatch(t, expected, insert(intervals, newInterval))
}
