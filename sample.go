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

// ErrSampleTooSmall is returned when the provided data sample set is too small
// for a computation.
//
// ErrInvalidConfidenceLevel is returned when the confidence level
// passed to MeanConfidenceIntervals is not in the valid range.
var (
	ErrSampleTooSmall         = errors.New("too few sample points")
	ErrInvalidConfidenceLevel = errors.New("invalid confidence level, 0<confidence<1)")
)

// Sample zero values are not safe, use the New function to initialize Sample
// types.
type Sample struct {
	data []float64
	// memoization:
	se *float64
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
		return nil, ErrSampleTooSmall
	}
	sample := new(Sample)
	sample.data = make([]float64, len(data))
	copy(sample.data, data)
	return sample, nil
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

// Mean computes the sample mean of a population sample.
//
// If the sample size is less than 1, it returns ErrSampleTooSmall
func Mean(data []float64) (float64, error) {
	if len(data) < 1 {
		return 0.0, ErrSampleTooSmall
	}

	return sum(data) / float64(len(data)), nil
}

func sum(s []float64) float64 {
	sum := 0.0
	for i := range s {
		sum += s[i]
	}

	return sum
}

// StandardDeviation computes the sample-based unbiased estimation of the
// standard deviation of a population.
//
// This implementation uses the Bessel's correction and is downward biassed as
// per Jensen's inequality.
//
// It is calculated as sqrt(1/(N-1) sum_i_N(x_i-mean(x))).
//
// If the sample size is less than 2, it returns ErrSampleTooSmall
func StandardDeviation(data []float64) (float64, error) {
	if len(data) < 2 {
		return 0.0, ErrSampleTooSmall
	}

	mean, _ := Mean(data)

	// TODO: this sum can be done concurrently
	sum := 0.0
	var diff float64
	for _, samplePoint := range data {
		diff = samplePoint - mean
		sum += diff * diff
	}

	return math.Sqrt(sum / float64(len(data)-1)), nil
}

// StandardError returns the standard deviation of the sampling
// distribution of the mean, also known as the "Standard Error of the
// Mean".
func (s *Sample) StandardErrorObj() float64 {
	if s.se != nil {
		return *s.se
	}
	s.se = new(float64)

	sd, _ := StandardDeviation(s.data)

	*s.se = sd / math.Sqrt(float64(len(s.data)))

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

	m, _ := Mean(s.data)

	return [2]float64{m - margin, m + margin}, nil
}
