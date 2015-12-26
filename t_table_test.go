package sample

import (
	"math"
	"testing"
)

func TestEqualOrLowerDegree(t *testing.T) {
	for i, f := range [...]struct {
		in  int64
		out int
	}{
		{1, 0},
		{2, 1},
		{10, 9},
		{39, 29},
		{73, 32},
		{545, 35},
		{math.MaxInt64, 36},
	} {
		got := equalOrLowerDegreeIndex(f.in)
		if fDegrees[got] != fDegrees[f.out] {
			t.Errorf("%d) in=%d, out=%d, got=%d", i, f.in, f.out, got)
		}
	}
}
