package sample

import (
	"math"
	"reflect"
	"testing"
)

const tolerance = 1e-3

func equals(a, b, tolerance float64) bool {
	if math.Abs(a-b) < tolerance {
		return true
	}
	return false
}

func TestNewErrors(t *testing.T) {
	for i, f := range [...][]float64{
		{},
		{1.0},
	} {
		s, err := New(f)
		if s != nil {
			t.Fatalf("%d) in=%v, should have returned nil, but got: %v", i, f, s)
		}
		if err != ErrBesselNeedsTwo {
			t.Fatalf("%d) in=%v, should have returned ErrBesselNeedsTwo", i, f)
		}
	}
}

func TestNew(t *testing.T) {
	for i, f := range [...][]float64{
		{1.0, 1},
		{1.0, 2, -3.1415952},
		{1.0, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
	} {
		s, err := New(f)
		if err != nil {
			t.Fatalf("%d) in=%v, returned error: %v", i, f, err)
		}
		if !reflect.DeepEqual(f, s.data) {
			t.Errorf("%d) in=%v, expected=%[2]v, got=%v", i, f, s.data)
		}
	}
}

func TestMean(t *testing.T) {
	for i, f := range [...]struct {
		in  []float64
		out float64
	}{
		{in: []float64{-1.0, -1.0}, out: -1.0},
		{in: []float64{1, 2, 3, 4, 5, 6}, out: 3.5},
		{in: []float64{0, 1, -1, 2, -2}, out: 0.0},
		{in: []float64{3.005, 3.005, 3.005, 3.005, 3.005, 3.005002, 3.004998}, out: 3.005},
	} {
		s, err := New(f.in)
		if err != nil {
			t.Fatalf("%d) in=%v, New returned error: %v", i, f.in, err)
		}
		got := s.Mean()
		if !equals(got, f.out, tolerance) {
			t.Errorf("%d) in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
	}
}

func TestMeanAlreadyComputed(t *testing.T) {
	f := []float64{1.0, 1}
	s, err := New(f)
	if err != nil {
		t.Fatalf("in=%v, New returned error: %v", f, err)
	}
	first := s.Mean()
	second := s.Mean()
	if !equals(first, second, tolerance) {
		t.Errorf("precomputed mean and new mean differs: sample=%v, first mean=%f, second mean=%f",
			f, first, second)
	}
}

func TestStandardDeviation(t *testing.T) {
	for i, f := range [...]struct {
		in  []float64
		out float64
	}{
		{in: []float64{1.0, 1.0}, out: 0.0},
		{in: []float64{1.0, 2.0}, out: 0.707},
		{in: []float64{1.0, 2, 3, 4, 5, 6}, out: 1.870},
		{in: []float64{-2.0, -1, 0, 1, 2, 3}, out: 1.870},
	} {
		s, err := New(f.in)
		if err != nil {
			t.Fatalf("%d) in=%v, New returned error: %v", i, f.in, err)
		}
		got := s.StandardDeviation()
		if !equals(got, f.out, tolerance) {
			t.Errorf("%d) in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
	}
}

func TestStandardDeviationAlreadyComputed(t *testing.T) {
	f := []float64{1.0, 1}
	s, err := New(f)
	if err != nil {
		t.Fatalf("in=%v, New returned error: %v", f, err)
	}
	first := s.StandardDeviation()
	second := s.StandardDeviation()
	if !equals(first, second, tolerance) {
		t.Errorf("precomputed standard deviation and new one differs: sample=%v, first sd=%f, second sd=%f",
			f, first, second)
	}
}

func TestStandardError(t *testing.T) {
	for i, f := range [...]struct {
		in  []float64
		out float64
	}{
		{in: []float64{1.0, 1.0}, out: 0.0},
		{in: []float64{1.0, 2.0}, out: 0.5},
		{in: []float64{1.0, 2, 3, 4, 5, 6}, out: 0.763},
		{in: []float64{-2.0, -1, 0, 1, 2, 3}, out: 0.763},
	} {
		s, err := New(f.in)
		if err != nil {
			t.Fatalf("%d) in=%v, New returned error: %v", i, f.in, err)
		}
		got := s.StandardError()
		if !equals(got, f.out, tolerance) {
			t.Errorf("%d) in=%v, out=%f, got=%f",
				i, f.in, f.out, got)
		}
	}
}

func TestStandardErrorAlreadyComputed(t *testing.T) {
	f := []float64{1.0, 1}
	s, err := New(f)
	if err != nil {
		t.Fatalf("in=%v, New returned error: %v", f, err)
	}
	first := s.StandardError()
	second := s.StandardError()
	if !equals(first, second, tolerance) {
		t.Errorf("precomputed standard error and new one differs: sample=%v, first sd=%f, second sd=%f",
			f, first, second)
	}
}
