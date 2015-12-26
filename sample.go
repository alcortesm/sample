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

// ErrBesselNeedsTwo is returned whenever the sample length is less than 2
// sample points, as required by Bessel's correction.
var (
	ErrBesselNeedsTwo = errors.New("Bessel's correction needs at least two sample points")
)

// Sample zero values are not safe, use the New function to initialize Sample
// types.
type Sample struct {
	data []float64
	// memoization:
	mean *float64
	sd   *float64
}

// New returns an Sample value initialized with a *copy* of its parameter and a
// nil error on success. It returns nil and ErrBesselNeedsTwo if the sample
// length is smaller than 2.
//
// A copy of the data is used internally to protect it from future
// modifications, allowing for memoizaton of already computed statistical
// values.
func New(data []float64) (*Sample, error) {
	if len(data) < 2 {
		return nil, ErrBesselNeedsTwo
	}
	sample := new(Sample)
	sample.data = make([]float64, len(data))
	copy(sample.data, data)
	return sample, nil
}

// Mean computes the sample mean of a population sample or returns its
// previously computed value.
func (s *Sample) Mean() float64 {
	if s.mean != nil {
		return *s.mean
	}
	s.mean = new(float64)

	sum := 0.0
	for i := range s.data {
		sum += s.data[i]
	}
	*s.mean = sum / float64(len(s.data))

	return *s.mean
}

// StandardDeviation computes the sample-based unbiased estimation of the
// standard deviation of a population or returns its previously computed value.
//
// This implementation uses the Bessel's correction and is downward biassed as
// per Jensen's inequality.
//
// It is calculated as sqrt(1/(N-1) sum_i_N(x_i-mean(x))).
func (s *Sample) StandardDeviation() float64 {
	if s.sd != nil {
		return *s.sd
	}
	s.sd = new(float64)

	s.Mean()

	sum := 0.0
	var diff float64
	for _, sp := range s.data {
		diff = sp - *s.mean
		sum += diff * diff
	}
	*s.sd = math.Sqrt(sum / float64(len(s.data)-1))

	return *s.sd
}
