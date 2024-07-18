
package array_string

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCanPlaceFlowers(t *testing.T) {
	flowerbed := []int{1,0,0,0,1}
	n := 2
	result := canPlaceFlowers(flowerbed,n)
	expected := false
	assert.Equal(t,expected,result)
}