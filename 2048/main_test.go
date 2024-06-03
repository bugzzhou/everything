package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveUp(t *testing.T) {
	inMap := [4][4]int{
		{0, 0, 0, 0},
		{1, 0, 0, 1},
		{0, 1, 0, 0},
		{1, 0, 0, 1},
	}
	expect := [4][4]int{
		{1, 1, 0, 1},
		{1, 0, 0, 1},
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}
	actual := inMap
	moveUp(&actual)

	assert.Equal(t, expect, actual)
}
