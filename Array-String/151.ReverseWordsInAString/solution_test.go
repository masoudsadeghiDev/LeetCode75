package array_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	result := reverseWords("the sky is blue")
	expected := "blue is sky the"
	assert.Equal(t,expected,result)
	
}