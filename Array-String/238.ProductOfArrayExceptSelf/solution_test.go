package array_string

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelf(t *testing.T) {
	nums:= []int{-1,1,0,-3,3}
	result := productExceptSelf(nums)
	expected := []int{0,0,9,0,0}
	assert.Equal(t,expected,result)
}