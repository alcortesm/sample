package sample

import (
	"math"
	"testing"
)

const tolerance = 1e-3

func equals(a, b float64) bool {
	if math.Abs(a-b) < tolerance {
		return true
	}
	return false
}

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
		if !equals(got, f.out) {
			t.Errorf("%d) in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
	}

	got, err := Mean([]float64{})
	if !math.IsNaN(got) {
		t.Errorf("empty slice) got %f, NaN was expected", got)
	}
	if err != ErrEmptyInputSample {
		t.Errorf("empty slice) got no ErrEmptyInputSample")
	}
}

func TestStandardDeviation(t *testing.T) {
	for i, f := range [...]struct {
		in  []float64
		out float64
	}{
		{
			in:  []float64{1.0, 1.0},
			out: 0.0,
		},
		{
			in:  []float64{1.0, 2.0},
			out: 0.707,
		},
		{
			in:  []float64{1.0, 2, 3, 4, 5, 6},
			out: 1.870,
		},
		{
			in:  []float64{-2.0, -1, 0, 1, 2, 3},
			out: 1.870,
		},
	} {
		got, err := StandardDeviation(f.in, nil)
		if err != nil {
			t.Errorf("%d) in=%v, returned error: %v", i, f.in, err)
		}
		if !equals(got, f.out) {
			t.Errorf("%d) in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
		// the same tests, but providing the mean
		mean, err := Mean(f.in)
		if err != nil {
			t.Errorf("%d) [with mean] unexpected error calculating the mean: %s", i, err)
		}
		got, err = StandardDeviation(f.in, &mean)
		if err != nil {
			t.Errorf("%d) [with mean] in=%v, returned error: %v", i, f.in, err)
		}
		if !equals(got, f.out) {
			t.Errorf("%d) [with mean] in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
	}

	// tests for empty sample
	got, err := StandardDeviation([]float64{}, nil)
	if !math.IsNaN(got) {
		t.Errorf("empty slice) got %f, NaN was expected", got)
	}
	if err != ErrEmptyInputSample {
		t.Errorf("empty slice) returned no ErrEmptyInputSample")
	}
}
