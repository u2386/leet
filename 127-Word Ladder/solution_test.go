package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestStringSlice(t *testing.T) {
	str := "hello"
	assert.Equal(t, "ell", str[1:4])
}

func TestStringPattern(t *testing.T) {
	contains := func (container []string, target string) bool {
		for _, s := range container {
			if (s == target) {
				return true
			}
		}
		return false
	}

	str := "hello"
	patterns := generatePattern(str)
	assert.Equal(t, 5, len(patterns))
	assert.True(t, contains(patterns, "*ello"))
	assert.True(t, contains(patterns, "h*llo"))
	assert.True(t, contains(patterns, "he*lo"))
	assert.True(t, contains(patterns, "hel*o"))
	assert.True(t, contains(patterns, "hell*"))
}
