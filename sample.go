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
	ErrSampleTooSmall    = errors.New("too few sample points")
	ErrInvalidConfidence = errors.New("invalid confidence level, 0 < confidence < 1)")
)

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
//
// If the sample size is less than 2, it returns ErrSampleTooSmall
func StandardError(data []float64) (float64, error) {
	sd, err := StandardDeviation(data)
	if err != nil {
		return 0.0, err
	}

	return sd / math.Sqrt(float64(len(data))), nil
}

// MeanConfidenceIntervals assumes the sample points are from a Normal
// distribution of unknown variance and calculates the confidence
// intervals of the mean of the distribution for the given confidence
// level.
//
// If the sample size is less than 2, it returns ErrSampleTooSmall.
//
// If the confidence value is not in the ]0, 1[ it returns
// ErrInvalidConfidence.
func MeanConfidenceIntervals(data []float64, confidence float64) ([2]float64,
	error) {

	if invalidConfidence(confidence) {
		return [2]float64{}, ErrInvalidConfidence
	}

	se, err := StandardError(data)
	if err != nil {
		return [2]float64{}, err
	}

	dimension := int64(len(data) - 1)
	tinv, err := studentTwoSidedCriticalValue(dimension, confidence)
	if err != nil {
		return [2]float64{}, err
	}
	margin := tinv * se

	mean, err := Mean(data)
	if err != nil {
		return [2]float64{}, err
	}

	return [2]float64{mean - margin, mean + margin}, nil
}

func invalidConfidence(confidence float64) bool {
	return confidence <= 0.0 || confidence >= 1.0
}
