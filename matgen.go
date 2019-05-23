package matgen

import (
	"math/rand"
	"time"

	"gonum.org/v1/gonum/mat"
)

type Generator struct {
	Rank    int
	Rows    int
	Columns int
}

type Option func(*Generator)

func Rank(r int) Option {
	return func(g *Generator) {
		g.Rank = r
	}
}

func Rows(r int) Option {
	return func(g *Generator) {
		if r <= 0 {
			panic("rows should be positive int")
		}
		g.Rows = r
	}
}

func Columns(c int) Option {
	return func(g *Generator) {
		if c <= 0 {
			panic("columns should be positive int")
		}
		g.Columns = c
	}
}

func New(opts ...Option) mat.Matrix {
	g := &Generator{}
	for _, opt := range opts {
		opt(g)
	}
	if g.Rank > g.Rows || g.Rank > g.Columns {
		panic("rank beyonds size of matrix")
	}
	rand.Seed(time.Now().UnixNano())
	var v []float64
	r, c := g.Rows, g.Columns
	if g.Rows > g.Columns {
		r, c = g.Columns, g.Rows
	}
	var first []float64
	for i := 0; i < r; i++ {
		if i > 0 && i < g.Rank {
			coef := rand.Int()
			if rand.Int()%2 == 0 {
				coef = -coef
			}
			for _, e := range first {
				v = append(v, e/float64(coef))
			}
		} else {
			for j := 0; j < c; j++ {
				e := (rand.Float64() - 0.5) * float64(rand.Int())
				v = append(v, e)
				if i == 0 {
					first = append(first, e)
				}
			}
		}
	}
	if g.Rows > g.Columns {
		return mat.NewDense(g.Rows, g.Columns, v).T()
	}
	return mat.NewDense(g.Rows, g.Columns, v)
}
