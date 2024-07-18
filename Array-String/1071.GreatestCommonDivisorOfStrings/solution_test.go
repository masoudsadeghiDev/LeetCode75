
package array_string

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGcdOfStrings(t *testing.T) {
	str1 := "ABABAB"
	str2 := "ABAB"
	result := gcdOfStrings(str1, str2)
	expected := "AB"
	assert.Equal(t,expected,result)
}