package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddNumber(t *testing.T) {

	add := AddNumber(10, 5)

	// assert equality
	assert.Equal(t, 15, add, "they should be equal")
}

func TestMinNumber(t *testing.T) {

	min := MinNumber(10, 5)

	// assert equality
	assert.Equal(t, 5, min, "they should be equal")
}
