package sample

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

const (
	causeEmptySlice = "empty slice"
	causeTooBig     = "too big"
	causeTooSmall   = "too small"
)

func TestIndexOfEqualOrClosestLower(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		s        []int64
		n        int64
		errCause string
		out      int
	}{
		{[]int64{}, 1, causeEmptySlice, 0},
		{[]int64{1}, 0, causeTooBig, 0},
		{[]int64{1}, 1, "", 0},
		{[]int64{1}, 2, "", 0},
		{[]int64{1, 3}, 0, causeTooBig, 0},
		{[]int64{1, 3}, 1, "", 0},
		{[]int64{1, 3}, 2, "", 0},
		{[]int64{1, 3}, 3, "", 1},
		{[]int64{1, 3}, 4, "", 1},
		{[]int64{1, 3, 5}, 0, causeTooBig, 0},
		{[]int64{1, 3, 5}, 1, "", 0},
		{[]int64{1, 3, 5}, 2, "", 0},
		{[]int64{1, 3, 5}, 3, "", 1},
		{[]int64{1, 3, 5}, 4, "", 1},
		{[]int64{1, 3, 5}, 5, "", 2},
		{[]int64{1, 3, 5}, 6, "", 2},
		{[]int64{1, 3, 5, 7}, 0, causeTooBig, 0},
		{[]int64{1, 3, 5, 7}, 1, "", 0},
		{[]int64{1, 3, 5, 7}, 2, "", 0},
		{[]int64{1, 3, 5, 7}, 3, "", 1},
		{[]int64{1, 3, 5, 7}, 4, "", 1},
		{[]int64{1, 3, 5, 7}, 5, "", 2},
		{[]int64{1, 3, 5, 7}, 6, "", 2},
		{[]int64{1, 3, 5, 7}, 7, "", 3},
		{[]int64{1, 3, 5, 7}, 8, "", 3},
		{[]int64{1, 3, 5, 7, 9}, 0, causeTooBig, 0},
		{[]int64{1, 3, 5, 7, 9}, 1, "", 0},
		{[]int64{1, 3, 5, 7, 9}, 2, "", 0},
		{[]int64{1, 3, 5, 7, 9}, 3, "", 1},
		{[]int64{1, 3, 5, 7, 9}, 4, "", 1},
		{[]int64{1, 3, 5, 7, 9}, 5, "", 2},
		{[]int64{1, 3, 5, 7, 9}, 6, "", 2},
		{[]int64{1, 3, 5, 7, 9}, 7, "", 3},
		{[]int64{1, 3, 5, 7, 9}, 8, "", 3},
		{[]int64{1, 3, 5, 7, 9}, 9, "", 4},
		{[]int64{1, 3, 5, 7, 9}, 10, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11}, 0, causeTooBig, 0},
		{[]int64{1, 3, 5, 7, 9, 11}, 1, "", 0},
		{[]int64{1, 3, 5, 7, 9, 11}, 2, "", 0},
		{[]int64{1, 3, 5, 7, 9, 11}, 3, "", 1},
		{[]int64{1, 3, 5, 7, 9, 11}, 4, "", 1},
		{[]int64{1, 3, 5, 7, 9, 11}, 5, "", 2},
		{[]int64{1, 3, 5, 7, 9, 11}, 6, "", 2},
		{[]int64{1, 3, 5, 7, 9, 11}, 7, "", 3},
		{[]int64{1, 3, 5, 7, 9, 11}, 8, "", 3},
		{[]int64{1, 3, 5, 7, 9, 11}, 9, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11}, 10, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11}, 11, "", 5},
		{[]int64{1, 3, 5, 7, 9, 11}, 12, "", 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 0, causeTooBig, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 1, "", 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 2, "", 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 3, "", 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 4, "", 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 5, "", 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 6, "", 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 7, "", 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 8, "", 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 9, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 10, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 11, "", 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 12, "", 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 13, "", 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 14, "", 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 15, "", 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 16, "", 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 17, "", 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 18, "", 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 19, "", 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 20, "", 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 21, "", 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 22, "", 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 0, causeTooBig, 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 1, "", 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 2, "", 0},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 3, "", 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 4, "", 1},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 5, "", 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 6, "", 2},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 7, "", 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 8, "", 3},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 9, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 10, "", 4},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 11, "", 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 12, "", 5},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 13, "", 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 14, "", 6},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 15, "", 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 16, "", 7},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 17, "", 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 18, "", 8},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 19, "", 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 20, "", 9},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 21, "", 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 22, "", 10},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 23, "", 11},
		{[]int64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 24, "", 11},
	} {
		test := test
		description := fmt.Sprintf("s=%v n=%v", test.s, test.n)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := indexOfEqualOrClosestLower(test.n, test.s)
			if test.errCause != "" {
				if err == nil {
					t.Fatalf("unexpected success, got %d", got)
				}
				if !strings.Contains(err.Error(), test.errCause) {
					t.Errorf("cannot find %q in %q", test.errCause, err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if test.s[got] != test.s[test.out] {
					t.Errorf("want %d, got %d", test.out, got)
				}
			}
		})
	}
}

func TestIndexOfEqualOrClosestHigher(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		s        []float64
		n        float64
		errCause string
		out      int
	}{
		{[]float64{}, 1, causeEmptySlice, 0},
		{[]float64{1}, 0, "", 0},
		{[]float64{1}, 1, "", 0},
		{[]float64{1}, 2, causeTooSmall, 0},
		{[]float64{1, 3}, 0, "", 0},
		{[]float64{1, 3}, 1, "", 0},
		{[]float64{1, 3}, 2, "", 1},
		{[]float64{1, 3}, 3, "", 1},
		{[]float64{1, 3}, 4, causeTooSmall, 1},
		{[]float64{1, 3, 5}, 0, "", 0},
		{[]float64{1, 3, 5}, 1, "", 0},
		{[]float64{1, 3, 5}, 2, "", 1},
		{[]float64{1, 3, 5}, 3, "", 1},
		{[]float64{1, 3, 5}, 4, "", 2},
		{[]float64{1, 3, 5}, 5, "", 2},
		{[]float64{1, 3, 5}, 6, causeTooSmall, 2},
		{[]float64{1, 3, 5, 7}, 0, "", 0},
		{[]float64{1, 3, 5, 7}, 1, "", 0},
		{[]float64{1, 3, 5, 7}, 2, "", 1},
		{[]float64{1, 3, 5, 7}, 3, "", 1},
		{[]float64{1, 3, 5, 7}, 4, "", 2},
		{[]float64{1, 3, 5, 7}, 5, "", 2},
		{[]float64{1, 3, 5, 7}, 6, "", 3},
		{[]float64{1, 3, 5, 7}, 7, "", 3},
		{[]float64{1, 3, 5, 7}, 8, causeTooSmall, 3},
		{[]float64{1, 3, 5, 7, 9}, 0, "", 0},
		{[]float64{1, 3, 5, 7, 9}, 1, "", 0},
		{[]float64{1, 3, 5, 7, 9}, 2, "", 1},
		{[]float64{1, 3, 5, 7, 9}, 3, "", 1},
		{[]float64{1, 3, 5, 7, 9}, 4, "", 2},
		{[]float64{1, 3, 5, 7, 9}, 5, "", 2},
		{[]float64{1, 3, 5, 7, 9}, 6, "", 3},
		{[]float64{1, 3, 5, 7, 9}, 7, "", 3},
		{[]float64{1, 3, 5, 7, 9}, 8, "", 4},
		{[]float64{1, 3, 5, 7, 9}, 9, "", 4},
		{[]float64{1, 3, 5, 7, 9}, 10, causeTooSmall, 4},
		{[]float64{1, 3, 5, 7, 9, 11}, 0, "", 0},
		{[]float64{1, 3, 5, 7, 9, 11}, 1, "", 0},
		{[]float64{1, 3, 5, 7, 9, 11}, 2, "", 1},
		{[]float64{1, 3, 5, 7, 9, 11}, 3, "", 1},
		{[]float64{1, 3, 5, 7, 9, 11}, 4, "", 2},
		{[]float64{1, 3, 5, 7, 9, 11}, 5, "", 2},
		{[]float64{1, 3, 5, 7, 9, 11}, 6, "", 3},
		{[]float64{1, 3, 5, 7, 9, 11}, 7, "", 3},
		{[]float64{1, 3, 5, 7, 9, 11}, 8, "", 4},
		{[]float64{1, 3, 5, 7, 9, 11}, 9, "", 4},
		{[]float64{1, 3, 5, 7, 9, 11}, 10, "", 5},
		{[]float64{1, 3, 5, 7, 9, 11}, 11, "", 5},
		{[]float64{1, 3, 5, 7, 9, 11}, 12, causeTooSmall, 5},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 0, "", 0},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 1, "", 0},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 2, "", 1},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 3, "", 1},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 4, "", 2},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 5, "", 2},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 6, "", 3},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 7, "", 3},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 8, "", 4},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 9, "", 4},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 10, "", 5},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 11, "", 5},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 12, "", 6},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 13, "", 6},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 14, "", 7},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 15, "", 7},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 16, "", 8},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 17, "", 8},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 18, "", 9},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 19, "", 9},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 20, "", 10},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 21, "", 10},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21}, 22, causeTooSmall, 10},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 0, "", 0},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 1, "", 0},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 2, "", 1},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 3, "", 1},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 4, "", 2},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 5, "", 2},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 6, "", 3},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 7, "", 3},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 8, "", 4},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 9, "", 4},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 10, "", 5},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 11, "", 5},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 12, "", 6},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 13, "", 6},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 14, "", 7},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 15, "", 7},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 16, "", 8},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 17, "", 8},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 18, "", 9},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 19, "", 9},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 20, "", 10},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 21, "", 10},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 22, "", 11},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 23, "", 11},
		{[]float64{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23}, 24, causeTooSmall, 11},
	} {
		test := test
		description := fmt.Sprintf("s=%v n=%v", test.s, test.n)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := indexOfEqualOrClosestHigher(test.n, test.s)
			if test.errCause != "" {
				if err == nil {
					t.Fatalf("unexpected success, got %d", got)
				}
				if !strings.Contains(err.Error(), test.errCause) {
					t.Errorf("cannot find %q in %q", test.errCause, err)
				}
			} else {
				if err != nil {
					t.Fatal(err)
				}
				if test.s[got] != test.s[test.out] {
					t.Errorf("want %d, got %d", test.out, got)
				}
			}
		})
	}
}

func TestStudentTwoSidedCriticalValue(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		degree     int64
		confidence float64
		want       float64
	}{
		{degree: 1, confidence: 0.50, want: 1.000},
		{degree: 2, confidence: 0.50, want: 0.816},
		{degree: math.MaxInt64, confidence: 0.50, want: 0.674},
		{degree: 10, confidence: 0.95, want: 2.228},
		{degree: 20, confidence: 0.95, want: 2.086},
		{degree: math.MaxInt64, confidence: 0.95, want: 1.960},
		{degree: 45, confidence: 0.96, want: 2.423}, // d=40, c=0.98
		{degree: 4, confidence: 0.95, want: 2.776},
	} {
		test := test
		description := fmt.Sprintf("degree=%d confidence=%f",
			test.degree, test.confidence)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := studentTwoSidedCriticalValue(test.degree, test.confidence)
			if err != nil {
				t.Fatal(err)
			}
			if !equals(test.want, got, 1e-5) {
				t.Errorf("want %f, got %f", test.want, got)
			}
		})
	}
}

func TestStudentTwoSidedCriticalValueErrorsInvalidConfidence(t *testing.T) {
	t.Parallel()
	for _, confidence := range []float64{
		-0.1, 0.0, 1.0, 1.1,
	} {
		confidence := confidence
		description := fmt.Sprint(confidence)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			_, err := studentTwoSidedCriticalValue(10, confidence)
			if err == nil {
				t.Fatal("unexpected success")
			}
			if err != ErrInvalidConfidence {
				t.Errorf("want: %q, got: %q", ErrInvalidConfidence, err)
			}
		})
	}
}

func TestStudentTwoSidedCriticalValueErrorLowFreedomDegree(t *testing.T) {
	t.Parallel()
	if _, err := studentTwoSidedCriticalValue(0, 0.50); err == nil {
		t.Errorf("unexpected success")
	}
}
