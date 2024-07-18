package array_string

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestReverseVowels(t *testing.T) {
	result := reverseVowels("leetcode")
	expected := "leotcede"
	assert.Equal(t,expected,result)
}