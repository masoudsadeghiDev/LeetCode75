package array_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncreasingTriplet(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	result := increasingTriplet(nums)
	expected := true
	assert.Equal(t, expected, result)

	nums = []int{5, 4, 3, 2, 1}
	result = increasingTriplet(nums)
	expected = false
	assert.Equal(t, expected, result)

}
