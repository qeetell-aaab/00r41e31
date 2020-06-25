package r0r41e31

import (
	"fmt"
	"testing"
)


func TestError (t *testing.T) {
	e1 := Error_New ("Error 1")
	fmt.Println ("i1:", e1.Error ())

	e2 := Error_New ("Error 2", e1)
	fmt.Println ("i2:", e2.Unwrap ().Error ())

	e3 := Error_New ("Error 2", e1)
	fmt.Println ("i2:", e3.Unwrap ().Error ())
}
