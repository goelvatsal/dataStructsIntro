package reverseString

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRevStr(t *testing.T) {
	in := "Hello, I am Vatsal."
	expected := ".lastaV ma I ,olleH"
	actual := revStr(in)

	assert.Equal(t, expected, actual)
}
