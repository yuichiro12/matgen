package matgen

import (
	"fmt"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestNew(t *testing.T) {
	m := New(Rank(10), Rows(20), Columns(30))
	fa := mat.Formatted(m, mat.Squeeze())
	fmt.Printf("%v\n", fa)

}
