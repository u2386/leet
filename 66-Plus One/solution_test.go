package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test01(t *testing.T) {
	arr := []int{1, 2, 3}
	assert.Equal(t, []int{1, 2, 4}, plusOne(arr))
}

func Test02(t *testing.T) {
	arr := []int{9, 9, 9}
	assert.Equal(t, []int{1, 0, 0, 0}, plusOne(arr))
}
