package nono

import (
	"testing"
)

func TestGen(t *testing.T) {
	var a = Nonogram{}
	a.Gen("simple")
	a.Display()
}
