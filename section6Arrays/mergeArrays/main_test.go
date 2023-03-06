package mergeArrays

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeArrays(t *testing.T) {
	in := []int{0, 3, 4, 31}
	in2 := []int{4, 6, 30}
	expected := []int{0, 3, 4, 4, 6, 30, 31}

	actual := mergeArrays(in, in2)
	assert.Equal(t, expected, actual)
}
