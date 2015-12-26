/*
Package sample implements some useful functions to process samples from
statistical populations. The standard Go `float64` type is used in all
computations.
*/
package sample

import (
	"errors"
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
