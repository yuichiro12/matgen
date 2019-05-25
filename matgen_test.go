package matgen

import (
	"fmt"
	"math/rand"
	"testing"

	"gonum.org/v1/gonum/mat"
)

func TestNew(t *testing.T) {
	for i := 0; i < 100; i++ {
		rows := rand.Intn(100) + 1
		cols := rand.Intn(100) + 1
		r := rand.Intn(min(rows, cols)) + 1
		epsilon := 1e-16
		mr, fa := calculateRank(r, rows, cols, epsilon)
		if mr != r {
			fmt.Printf("singular value vector: %v\n", fa)
			t.Errorf("rows %v, cols %v, got %v, want %v", rows, cols, mr, r)
		}
	}
}

func calculateRank(rank, rows, cols int, epsilon float64) (int, fmt.Formatter) {
	m := New(Rank(rank), Rows(rows), Columns(cols))
	fa := mat.Formatted(m, mat.Squeeze())
	return MatrixRankWithDumpSV(m, epsilon), fa

}
