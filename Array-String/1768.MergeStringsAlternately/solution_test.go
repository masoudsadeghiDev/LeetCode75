package array_string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMergeAlternately(t *testing.T) {
	result := mergeAlternately("abcd", "pq")
	expected := "apbqcd"
	assert.Equal(t, expected, result)
}
