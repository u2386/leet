package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCase01(t *testing.T) {
	heights := []int{1, 2, 3, 4}
	expected := []int{2, 3, 3, 4}

	assert.Equal(t, expected, pourWater(heights, 2, 2))
}

func TestCase02(t *testing.T) {
	heights := []int{2, 1, 1, 2, 1, 2, 2}
	expected := []int{2, 2, 2, 3, 2, 2, 2}

	assert.Equal(t, expected, pourWater(heights, 4, 3))
}

func TestCase03(t *testing.T) {
	heights := []int{3, 1, 3}
	expected := []int{4, 4, 4}

	assert.Equal(t, expected, pourWater(heights, 5, 1))
}
