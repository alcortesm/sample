package sample

import (
	"math"
	"testing"
)

func TestMean(t *testing.T) {
	for i, f := range [...]struct {
		in  []float64
		out float64
	}{
		{
			in:  []float64{1.0},
			out: 1.0,
		},
		{
			in:  []float64{-1.0, -1.0},
			out: -1.0,
		},
		{
			in:  []float64{1, 2, 3, 4, 5, 6},
			out: 3.5,
		},
		{
			in:  []float64{0, 1, -1, 2, -2},
			out: 0.0,
		},
		{
			in: []float64{
				3.005,
				3.005,
				3.005,
				3.005,
				3.005,
				3.005002,
				3.004998,
			},
			out: 3.005,
		},
	} {
		got, err := Mean(f.in)
		if err != nil {
			t.Errorf("%d) in=%v, returned error: %v", i, f.in, err)
		}
		if got != f.out {
			t.Errorf("%d) in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
	}

	got, err := Mean([]float64{})
	if !math.IsNaN(got) {
		t.Errorf("empty slice) got %f, NaN was expected")
	}
	if err == nil {
		t.Errorf("empty slice) got no error")
	}
}
