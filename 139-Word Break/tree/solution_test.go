package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test01(t *testing.T) {
	s := "abcd"
	wd := []string{"a", "abc", "b", "cd"}

	assert.True(t, wordBreak(s, wd))
}

func Test02(t *testing.T) {
	s := "bb"
	wd := []string{"a", "b", "bbb", "bbb"}

	assert.True(t, wordBreak(s, wd))
}

func Test03(t *testing.T) {
	s := "leetcode"
	wd := []string{"leet", "code"}

	assert.True(t, wordBreak(s, wd))
}

func Test04(t *testing.T) {
	s := "catsandog"
	wd := []string{"cats", "dog", "sand", "and", "cat"}

	assert.False(t, wordBreak(s, wd))
}
