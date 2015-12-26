/*
Package sample implements some useful functions to process samples from
statistical populations. The standard Go `float64` type is used in all
computations.
*/
package sample

import (
	"fmt"
	"math"
)

// Mean returns the sample mean and a nil error on success.  If the
// input parameter has zero elements, it returns `math.NaN()` and an
// error.
func Mean(s []float64) (float64, error) {
	if len(s) == 0 {
		return math.NaN(), fmt.Errorf("empty sample")
	}
	sum := 0.0
	for _, sp := range s {
		sum += sp
	}
	return sum / float64(len(s)), nil
}
