package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemove(t *testing.T) {
	raw := []int{1, 2, 3, 4, 5}
	expect := []int{1, 3, 4, 5}
	actual := RemoveSliceByIndex(raw, 1)
	assert.Equal(t, expect, actual)
}
