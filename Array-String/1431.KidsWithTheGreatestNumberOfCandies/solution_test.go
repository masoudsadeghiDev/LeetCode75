package array_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKidsWithCandies(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	result := kidsWithCandies(candies, extraCandies)
	expected := []bool{true, true, true, false, true}
	assert.Equal(t, expected, result)
}
