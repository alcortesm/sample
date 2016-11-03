package sample

import (
	"errors"
	"fmt"
	"math"
)

var (
	errEmptySlice   = errors.New("empty slice")
	errDataTooBig   = errors.New("no lower data value found")
	errDataTooSmall = errors.New("no bigger data value found")
)

// Returns the 2-sided critical values of a Student-t distribution with
// 'd' degrees of freedom and a percentile of 'p'.
//
// The values are looked up in a table, using the lower closest
// approximation to 'd' and higer closest approximaiton of 'c' in the
// table (this is, it returns a conservative approximation for values
// not present in the table).
func studentTwoSidedCriticalValue(d int64, c float64) (float64, error) {
	type resultAndError struct {
		index int
		err   error
	}
	ch := make(chan *resultAndError)
	go func() {
		i, err := indexOfEqualOrClosestLower(d, degrees)
		ch <- &resultAndError{index: i, err: err}
	}()

	confidenceIndex, err := indexOfEqualOrClosestHigher(c, percentile)
	if err != nil {
		return 0.0, fmt.Errorf("cannot approximate confidence: %s", err)
	}

	degree := <-ch
	if degree.err != nil {
		return 0.0, fmt.Errorf("cannot approximate degrees of freedom: %s", degree.err)
	}

	return tTable[degree.index][confidenceIndex], nil
}

// Finds the index of integer 'n' in a sorted (ascending) slice 's'.
//
// If 'n' is not found in 's', it returns the index of the closest
// integer to 'n' that is smaller than 'n'.
//
// If all integers in 's' are bigger than 'n', it returns errDataTooBig.
//
// If 's' is empty, it returns errEmptySlice.
func indexOfEqualOrClosestLower(n int64, s []int64) (i int, err error) {
	if len(s) == 0 {
		return 0, errEmptySlice
	}

	if n < s[0] {
		return 0, errDataTooBig
	}

	if n > s[len(s)-1] {
		return len(s) - 1, nil
	}

	b := 0
	e := len(s) - 1
	var m int
	for {
		if e-b < 2 {
			if s[e] == n {
				return e, nil
			}
			return b, nil
		}

		m = b + (e-b)/2
		if n == s[m] {
			return m, nil
		}
		if n < s[m] {
			e = m
		} else {
			b = m
		}
	}
}

// Finds the index of integer 'n' in a sorted (ascending) slice 's'.
//
// If 'n' is not found in 's', it returns the index of the closest
// integer to 'n' that is smaller than 'n'.
//
// If all integers in 's' are bigger than 'n', it returns errDataTooBig.
//
// If 's' is empty, it returns errEmptySlice.
func indexOfEqualOrClosestHigher(n float64, s []float64) (i int, err error) {
	if len(s) == 0 {
		return 0, errEmptySlice
	}

	if n < s[0] {
		return 0, nil
	}

	if n > s[len(s)-1] {
		return 0, errDataTooSmall
	}

	b := 0
	e := len(s) - 1
	var m int
	for {
		if e-b < 2 {
			if s[b] == n {
				return b, nil
			}
			return e, nil
		}

		m = b + (e-b)/2
		if n == s[m] {
			return m, nil
		}
		if n < s[m] {
			e = m
		} else {
			b = m
		}
	}
}

// from https://en.wikipedia.org/wiki/Student%27s_t-distribution#Table_of_selected_values
var degrees = []int64{
	1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
	21, 22, 23, 24, 25, 26, 27, 28, 29, 30,
	40, 50, 60, 80, 100, 120, math.MaxInt64,
}

var percentile = []float64{
	0.50, 0.60, 0.70, 0.80, 0.90, 0.95, 0.98, 0.99, 0.995, 0.998, 0.999,
}
var tTable = [][]float64{
	{1.000, 1.376, 1.963, 3.078, 6.314, 12.71, 31.82, 63.66, 127.3, 318.3, 636.6},
	{0.816, 1.080, 1.386, 1.886, 2.920, 4.303, 6.965, 9.925, 14.09, 22.33, 31.60},
	{0.765, 0.978, 1.250, 1.638, 2.353, 3.182, 4.541, 5.841, 7.453, 10.21, 12.92},
	{0.741, 0.941, 1.190, 1.533, 2.132, 2.776, 3.747, 4.604, 5.598, 7.173, 8.610},
	{0.727, 0.920, 1.156, 1.476, 2.015, 2.571, 3.365, 4.032, 4.773, 5.893, 6.869},
	{0.718, 0.906, 1.134, 1.440, 1.943, 2.447, 3.143, 3.707, 4.317, 5.208, 5.959},
	{0.711, 0.896, 1.119, 1.415, 1.895, 2.365, 2.998, 3.499, 4.029, 4.785, 5.408},
	{0.706, 0.889, 1.108, 1.397, 1.860, 2.306, 2.896, 3.355, 3.833, 4.501, 5.041},
	{0.703, 0.883, 1.100, 1.383, 1.833, 2.262, 2.821, 3.250, 3.690, 4.297, 4.781},
	{0.700, 0.879, 1.093, 1.372, 1.812, 2.228, 2.764, 3.169, 3.581, 4.144, 4.587},
	{0.697, 0.876, 1.088, 1.363, 1.796, 2.201, 2.718, 3.106, 3.497, 4.025, 4.437},
	{0.695, 0.873, 1.083, 1.356, 1.782, 2.179, 2.681, 3.055, 3.428, 3.930, 4.318},
	{0.694, 0.870, 1.079, 1.350, 1.771, 2.160, 2.650, 3.012, 3.372, 3.852, 4.221},
	{0.692, 0.868, 1.076, 1.345, 1.761, 2.145, 2.624, 2.977, 3.326, 3.787, 4.140},
	{0.691, 0.866, 1.074, 1.341, 1.753, 2.131, 2.602, 2.947, 3.286, 3.733, 4.073},
	{0.690, 0.865, 1.071, 1.337, 1.746, 2.120, 2.583, 2.921, 3.252, 3.686, 4.015},
	{0.689, 0.863, 1.069, 1.333, 1.740, 2.110, 2.567, 2.898, 3.222, 3.646, 3.965},
	{0.688, 0.862, 1.067, 1.330, 1.734, 2.101, 2.552, 2.878, 3.197, 3.610, 3.922},
	{0.688, 0.861, 1.066, 1.328, 1.729, 2.093, 2.539, 2.861, 3.174, 3.579, 3.883},
	{0.687, 0.860, 1.064, 1.325, 1.725, 2.086, 2.528, 2.845, 3.153, 3.552, 3.850},
	{0.686, 0.859, 1.063, 1.323, 1.721, 2.080, 2.518, 2.831, 3.135, 3.527, 3.819},
	{0.686, 0.858, 1.061, 1.321, 1.717, 2.074, 2.508, 2.819, 3.119, 3.505, 3.792},
	{0.685, 0.858, 1.060, 1.319, 1.714, 2.069, 2.500, 2.807, 3.104, 3.485, 3.767},
	{0.685, 0.857, 1.059, 1.318, 1.711, 2.064, 2.492, 2.797, 3.091, 3.467, 3.745},
	{0.684, 0.856, 1.058, 1.316, 1.708, 2.060, 2.485, 2.787, 3.078, 3.450, 3.725},
	{0.684, 0.856, 1.058, 1.315, 1.706, 2.056, 2.479, 2.779, 3.067, 3.435, 3.707},
	{0.684, 0.855, 1.057, 1.314, 1.703, 2.052, 2.473, 2.771, 3.057, 3.421, 3.690},
	{0.683, 0.855, 1.056, 1.313, 1.701, 2.048, 2.467, 2.763, 3.047, 3.408, 3.674},
	{0.683, 0.854, 1.055, 1.311, 1.699, 2.045, 2.462, 2.756, 3.038, 3.396, 3.659},
	{0.683, 0.854, 1.055, 1.310, 1.697, 2.042, 2.457, 2.750, 3.030, 3.385, 3.646},
	{0.681, 0.851, 1.050, 1.303, 1.684, 2.021, 2.423, 2.704, 2.971, 3.307, 3.551},
	{0.679, 0.849, 1.047, 1.299, 1.676, 2.009, 2.403, 2.678, 2.937, 3.261, 3.496},
	{0.679, 0.848, 1.045, 1.296, 1.671, 2.000, 2.390, 2.660, 2.915, 3.232, 3.460},
	{0.678, 0.846, 1.043, 1.292, 1.664, 1.990, 2.374, 2.639, 2.887, 3.195, 3.416},
	{0.677, 0.845, 1.042, 1.290, 1.660, 1.984, 2.364, 2.626, 2.871, 3.174, 3.390},
	{0.677, 0.845, 1.041, 1.289, 1.658, 1.980, 2.358, 2.617, 2.860, 3.160, 3.373},
	{0.674, 0.842, 1.036, 1.282, 1.645, 1.960, 2.326, 2.576, 2.807, 3.090, 3.291},
}
