/*
Package sample implements some useful functions to process samples from
statistical populations. The standard Go `float64` type is used in all
computations.
*/
package sample

import (
	"errors"
	"fmt"
	"math"
)

// ErrEmptyInputSample is returned whenever a function in this package receives
// an empty slice as the population sample to process.
var ErrEmptyInputSample = errors.New("empty input sample")

// Mean computes the sample mean of a population sample. It receives the
// population sample as a slice of float64 and returns the sample mean and a
// nil error on success.  If the input parameter has zero elements, it returns
// math.NaN() and ErrEmptyInputSample.
func Mean(s []float64) (float64, error) {
	if len(s) == 0 {
		return math.NaN(), ErrEmptyInputSample
	}
	sum := 0.0
	for _, sp := range s {
		sum += sp
	}
	return sum / float64(len(s)), nil
}

// StandardDeviation computes the sample-based unbiased estimation of the
// standard deviation of a population.
//
// It receives the population sample as a slice of float64 and (optionally) its
// mean. The function returns the estimate standard deviation of the population
// and a nil error.  If the input parameter has zero elements, it returns
// math.NaN() and ErrEmptyInputSample.
//
// If the provided mean is nil, it will be calculated by the function.
//
// This implementation uses the Bessel's correction and therefore needs at
// least a sample population of 2 sample points or more.
//
// This implementation is downward biassed as per Jensen's inequality.
//
// It is calculated as sqrt(1/(N-1) sum_i_N(x_i-mean(x))).
func StandardDeviation(s []float64, mean *float64) (float64, error) {
	if len(s) == 0 {
		return math.NaN(), ErrEmptyInputSample
	}
	if len(s) == 1 {
		return math.NaN(), fmt.Errorf("Bessel's correction needs at least two sample points")
	}
	if mean == nil {
		mean = new(float64)
		*mean, _ = Mean(s) // empty slice error already checked
	}
	sum := 0.0
	for _, sp := range s {
		diff := sp - *mean
		squared := diff * diff
		sum += squared
	}
	return math.Sqrt(sum / float64(len(s)-1)), nil
}
