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

// ErrBesselNeedsTwo is returned whenever the sample length is less than
// 2 sample points, as required by Bessel's correction.
//
// ErrInvalidConfidenceLevel is returned when the confidence level
// passed to MeanConfidenceIntervals is not in the valid range.
var (
	ErrBesselNeedsTwo         = errors.New("Bessel's correction needs at least two sample points")
	ErrInvalidConfidenceLevel = errors.New("invalid confidence level, 0<confidence<1)")
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

// New returns a Sample value initialized with a *copy* of its parameter and a
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

// Split splits a slice s into n sub-slices. The order of the elements
// in the resulting slices is such that concatenating them, in the same
// order they are returned, will result in the original slice.
//
// The number of elements in each sub-slice is len(s)/n in case n is a
// factor of len(s). If not, the first len(s)%n sub-slices will have
// one additional element.
//
// If the s is nil, the return value will be n slices of nil.
func split(s []float64, n int) [][]float64 {
	r := make([][]float64, n)
	if n == 0 {
		return r
	}
	quotient := len(s) / n
	remainder := len(s) % n
	begin := 0
	for i := 0; i < n; i++ {
		end := begin + quotient
		if i < remainder {
			end++
		}
		r[i] = s[begin:end]
		begin = end
	}
	return r
}

func sumConcurrent(s []float64) float64 {
	n := 4
	ch := make(chan float64, n)
	portions := split(s, n)
	for _, p := range portions {
		go func(data []float64) {
			ch <- sum(data)
		}(p)
	}

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += <-ch
	}

	return sum
}

// Mean computes the sample mean of a population sample or returns its
// previously computed value.
func (s *Sample) MeanObj() float64 {
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
func (s *Sample) StandardDeviationObj() float64 {
	if s.sd != nil {
		return *s.sd
	}
	s.sd = new(float64)

	s.MeanObj()

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
func (s *Sample) StandardErrorObj() float64 {
	if s.se != nil {
		return *s.se
	}
	s.se = new(float64)

	s.StandardDeviationObj()

	*s.se = *s.sd / math.Sqrt(float64(len(s.data)))

	return *s.se
}

// MeanConfidenceIntervals assumes the sample came from a Normal
// distribution of unknown variance and calculates the confidence
// intervals of the mean of the distribution for a confidence level of
// `c`.
func (s *Sample) MeanConfidenceIntervalsObj(c float64) ([2]float64, error) {
	if c <= 0.0 || c >= 1.0 {
		return [2]float64{}, ErrInvalidConfidenceLevel
	}

	dimension := int64(len(s.data) - 1)

	tinv, err := studentTwoSidedCriticalValue(dimension, c)
	if err != nil {
		return [2]float64{}, err
	}

	margin := tinv * s.StandardErrorObj()

	return [2]float64{s.MeanObj() - margin, s.MeanObj() + margin}, nil
}
