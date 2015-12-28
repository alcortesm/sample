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
//
// ErrNotAPercentile is returned when the confidence level for
// MeanConfidenceIntervals is not in the percentile range.
var (
	ErrBesselNeedsTwo     = errors.New("Bessel's correction needs at least two sample points")
	ErrNotAPercentile     = errors.New("number is not a percentile (0<=x<=1)")
	ErrNotPositiveInteger = errors.New("number is not a positive integer (0<n)")
)

// Sample zero values are not safe, use the New function to initialize Sample
// types.
type Sample struct {
	data []float64
	// memoization:
	mean *float64
	sd   *float64
	se   *float64
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

func sum(s []float64) float64 {
	sum := 0.0
	for i := range s {
		sum += s[i]
	}
	return sum
}

// Mean computes the sample mean of a population sample or returns its
// previously computed value.
func (s *Sample) Mean() float64 {
	if s.mean != nil {
		return *s.mean
	}
	s.mean = new(float64)

	*s.mean = sum(s.data) / float64(len(s.data))

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

	// TODO: this sum can be done concurrently
	sum := 0.0
	var diff float64
	for _, sp := range s.data {
		diff = sp - *s.mean
		sum += diff * diff
	}
	*s.sd = math.Sqrt(sum / float64(len(s.data)-1))

	return *s.sd
}

// StandardError returns the standard deviation of the sampling
// distribution of the mean, also known as the "Standard Error of the
// Mean".
func (s *Sample) StandardError() float64 {
	if s.se != nil {
		return *s.se
	}
	s.se = new(float64)

	s.StandardDeviation()

	*s.se = *s.sd / math.Sqrt(float64(len(s.data)))

	return *s.se
}
