package nono

import (
	"fmt"
	"testing"
)

func TestGen(t *testing.T) {
	var a = Nonogram{}
	a.Gen("normal")
	a.Display()
	fmt.Printf("%v\n", a.Check())
}
