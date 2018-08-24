[![GoDoc](https://godoc.org/github.com/alcortesm/sample?status.svg)](https://godoc.org/github.com/alcortesm/sample)
[![Build Status](https://travis-ci.org/alcortesm/sample.png)](https://travis-ci.org/alcortesm/sample)
[![codecov.io](https://codecov.io/github/alcortesm/sample/coverage.svg?branch=master)](https://codecov.io/github/alcortesm/sample?branch=master)

# Sample

Sample is a Go package to calculate useful values from samples of statistical
populations, including:

- the sample mean

- the sample-based unbiased estimation of the standard deviation of the
  population

- the standard error of the mean

- the confidence intervals of the mean, assuming the samples comes from a Normal
  distribution of unknown variance

The standard Go float64 type is used in all computations.

This package does *not* take advantage of multicore architectures.

## Import

```
import "github.com/alcortesm/sample"
```

## Examples

```Go
package main

import (
	"fmt"
	"log"

	"github.com/alcortesm/sample"
)

func main() {
	data := []float64{1.1, 0.9, 1.1, 1.3, 1.0}

	// ignoring errors for demonstration purposes
	mean, _ := sample.Mean(data)
	fmt.Println(mean) // 1.08

	sd, _ := sample.StandardDeviation(data)
	fmt.Println(sd) // 0.15

	se, _ := sample.StandardError(data)
	fmt.Println(se) // 0.07

	confidence := 0.95
	ci, _ := sample.MeanConfidenceIntervals(data, confidence)
	fmt.Printf("[%1.2f, %1.2f]\n", ci[0], ci[1]) // [0.90, 1.26]
}
```

## Author

- Alberto Cortés <alcortesm@gmail.com>.

## License

This project is licensed under the MIT License - see the
[LICENSE](LICENSE) file for details.

