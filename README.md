[![GoDoc](https://godoc.org/github.com/alcortesm/sample?status.svg)](https://godoc.org/github.com/alcortesm/sample)
[![Build Status](https://travis-ci.org/alcortesm/sample.png)](https://travis-ci.org/alcortesm/sample)
[![codecov.io](https://codecov.io/github/alcortesm/sample/coverage.svg?branch=master)](https://codecov.io/github/alcortesm/sample?branch=master)

# Sample

Sample is a Go package to calculate useful values from samples of statistical
populations, like:

- the sample mean

- the sample-based unbiased estimation of the standard deviation of the
  population

- the standard error of the mean

- the confidence intervals of the mean, assuming the samples comes from a Normal
  distribution of unknown variance

The standard Go float64 type is used in all computations.

This package does *not* take advantage of multicore architectures.

## Installation

```bash
go get github.com/alcortesm/sample
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

	mean, err := sample.Mean(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mean = %1.2f\n", mean) // 1.08

	sd, err := sample.StandardDeviation(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("std. deviation = %1.2f\n", sd) // 0.15

	se, err := sample.StandardError(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("std. error = %1.2f\n", se) // 0.07

	confidence := 0.95
	ci, err := sample.MeanConfidenceIntervals(data, confidence)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("mean conf. intervals (%2.0f%%) = [%1.2f, %1.2f]\n",
		100*confidence, ci[0], ci[1]) // [0.90, 1.26]
}
```

## Author

- Alberto Cortés <alcortesm@gmail.com>

## License

This project is licensed under the GPLv2 License - see the
[LICENSE.md](LICENSE.md) file for details

