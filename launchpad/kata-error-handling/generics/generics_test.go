package generics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumOfInts(t *testing.T) {
	assert.Equal(t, 7, Sum([]int{1, 2, 3, 1}))
}

func TestSumOfStrings(t *testing.T) {
	assert.Equal(t, "Hello World !", Sum([]string{"Hello", " World", " !"}))
}

func TestSumOfFloats64(t *testing.T) {
	assert.Equal(t, 34.05870001, Sum([]float64{3.4, 5.202, 1.00000001, 24.4567}))
}

func TestSumOfFloats32(t *testing.T) {
	assert.Equal(t, float32(34.0587), Sum([]float32{3.4, 5.202, 1.00000001, 24.4567}))
}