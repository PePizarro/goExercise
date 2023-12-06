package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSliceFunc(t *testing.T) {

	assert := assert.New(t)
	input := []float64{1.0, 2.0, 3.0}
	expected := []float64{1.0, 4.0, 9.0}

	out := newSlice(input)
	assert.ElementsMatch(out, expected, "The two slices should be the same")
}
