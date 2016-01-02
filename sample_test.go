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

func TestSumSamll(t *testing.T) {
	for i, f := range [...]struct {
		input    []float64
		expected float64
	}{
		{input: []float64{}, expected: 0.0},
		{input: []float64{3.1}, expected: 3.1},
		{input: []float64{3.1, 2.0}, expected: 5.1},
		{input: []float64{3.1, 2.0, 7.0}, expected: 12.1},
		{input: []float64{1.0, -1.0, 3.0}, expected: 3.0},
		{input: []float64{1.1, 2.0, 3.0, 4.0}, expected: 10.1},
	} {
		output := sum(f.input)
		if !equals(output, f.expected, tolerance) {
			t.Errorf("%d) input=%v, expected=%f, output=%f",
				i, f.input, f.expected, output)
		}
	}
}

func oneToN(n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = float64(i + 1)
	}
	return s
}

// this will tests the sum of big slices, using simple finite arithmetic
// progressions as the input data.
func TestSumBig(t *testing.T) {
	for i, n := range [...]int{
		10,
		100,
		1000,
		10000,
		100000,
	} {
		expected := float64(n*(1+n)) / 2
		input := oneToN(n)
		output := sum(input)
		if !equals(output, expected, tolerance) {
			t.Errorf("%d) n=%d, expected=%f, output=%f",
				i, n, expected, output)
		}
	}
}

var benchmarkResult float64 // avoid compiler optimization to elimitate tests

func benchmarkSum(n int, b *testing.B) {
	b.StopTimer()
	var r float64
	input := oneToN(n)
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		r = sum(input)
	}
	benchmarkResult = r
}

func BenchmarkSum2(b *testing.B) { benchmarkSum(10, b) }
func BenchmarkSum3(b *testing.B) { benchmarkSum(100, b) }
func BenchmarkSum4(b *testing.B) { benchmarkSum(1000, b) }
func BenchmarkSum5(b *testing.B) { benchmarkSum(10000, b) }
func BenchmarkSum6(b *testing.B) { benchmarkSum(100000, b) }

func TestSplit(t *testing.T) {
	for i, tt := range [...]struct {
		s        []float64
		n        int
		expected [][]float64
	}{
		{nil, 0, [][]float64{}},
		{[]float64{}, 0, [][]float64{}},
		{[]float64{1}, 0, [][]float64{}},
		{[]float64{1, 2}, 0, [][]float64{}},
		{[]float64{1, 2, 3}, 0, [][]float64{}},
		{[]float64{1, 2, 3, 4}, 0, [][]float64{}},
		{nil, 1, [][]float64{nil}}, // 6
		{[]float64{}, 1, [][]float64{{}}},
		{[]float64{1}, 1, [][]float64{{1}}},
		{[]float64{1, 2}, 1, [][]float64{{1, 2}}},
		{[]float64{1, 2, 3}, 1, [][]float64{{1, 2, 3}}},
		{[]float64{1, 2, 3, 4}, 1, [][]float64{{1, 2, 3, 4}}},
		{nil, 2, [][]float64{nil, nil}}, // 12
		{[]float64{},
			2, [][]float64{
				{},
				{}}},
		{[]float64{1},
			2, [][]float64{
				{1},
				{}}},
		{[]float64{1, 2},
			2, [][]float64{
				{1},
				{2}}},
		{[]float64{1, 2, 3},
			2, [][]float64{
				{1, 2},
				{3}}},
		{[]float64{1, 2, 3, 4},
			2, [][]float64{
				{1, 2},
				{3, 4}}},
		{[]float64{1, 2, 3, 4, 5},
			2, [][]float64{
				{1, 2, 3},
				{4, 5}}},
		{[]float64{1, 2, 3, 4, 5, 6},
			2, [][]float64{
				{1, 2, 3},
				{4, 5, 6}}},
		{[]float64{1, 2, 3, 4, 5, 6, 7},
			2, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7}}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8},
			2, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8}}},
		{nil, 3, [][]float64{nil, nil, nil}},
		{[]float64{},
			3, [][]float64{
				{},
				{},
				{}}},
		{[]float64{1},
			3, [][]float64{
				{1},
				{},
				{}}},
		{[]float64{1, 2},
			3, [][]float64{
				{1},
				{2},
				{}}},
		{[]float64{1, 2, 3},
			3, [][]float64{
				{1},
				{2},
				{3}}},
		{[]float64{1, 2, 3, 4},
			3, [][]float64{
				{1, 2},
				{3},
				{4}}},
		{[]float64{1, 2, 3, 4, 5},
			3, [][]float64{
				{1, 2},
				{3, 4},
				{5}}},
		{[]float64{1, 2, 3, 4, 5, 6},
			3, [][]float64{
				{1, 2},
				{3, 4},
				{5, 6}}},
		{[]float64{1, 2, 3, 4, 5, 6, 7},
			3, [][]float64{
				{1, 2, 3},
				{4, 5},
				{6, 7}}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8},
			3, [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8}}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			3, [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9}}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			3, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7},
				{8, 9, 10}}},
		{nil, 4, [][]float64{nil, nil, nil, nil}},
		{[]float64{},
			4, [][]float64{
				{},
				{},
				{},
				{}}},
		{[]float64{1},
			4, [][]float64{
				{1},
				{},
				{},
				{}}},
		{[]float64{1, 2},
			4, [][]float64{
				{1},
				{2},
				{},
				{}}},
		{[]float64{1, 2, 3},
			4, [][]float64{
				{1},
				{2},
				{3},
				{}}},
		{[]float64{1, 2, 3, 4},
			4, [][]float64{
				{1},
				{2},
				{3},
				{4}}},
		{[]float64{1, 2, 3, 4, 5},
			4, [][]float64{
				{1, 2},
				{3},
				{4},
				{5},
			}},
		{[]float64{1, 2, 3, 4, 5, 6},
			4, [][]float64{
				{1, 2},
				{3, 4},
				{5},
				{6},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7},
			4, [][]float64{
				{1, 2},
				{3, 4},
				{5, 6},
				{7},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8},
			4, [][]float64{
				{1, 2},
				{3, 4},
				{5, 6},
				{7, 8},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			4, [][]float64{
				{1, 2, 3},
				{4, 5},
				{6, 7},
				{8, 9},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			4, [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8},
				{9, 10},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
			4, [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12},
			4, [][]float64{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11, 12},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			4, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7},
				{8, 9, 10},
				{11, 12, 13},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14},
			4, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11},
				{12, 13, 14},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15},
			4, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15},
			}},
		{[]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16},
			4, [][]float64{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			}},
	} {
		output := split(tt.s, tt.n)
		if !reflect.DeepEqual(output, tt.expected) {
			t.Errorf("%d) s=%v, n=%d\nexpect=%v\noutput=%v",
				i, tt.s, tt.n, tt.expected, output)
		}
	}
}
