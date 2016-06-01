package sample

import (
	"math"
	"testing"
)

const tolerance = 1e-3

func equals(a, b, tolerance float64) bool {
	if math.Abs(a-b) < tolerance {
		return true
	}
	return false
}

func pairEquals(a, b [2]float64, tolerance float64) bool {
	for i := 0; i < len(a); i++ {
		if !equals(a[i], b[i], tolerance) {
			return false
		}
	}

	return true
}

func TestMean(t *testing.T) {
	for testNumber, testData := range [...]struct {
		input          []float64
		expectedError  error
		expectedResult float64
	}{
		{input: nil, expectedError: ErrSampleTooSmall},
		{input: []float64{}, expectedError: ErrSampleTooSmall},
		{input: []float64{12.}, expectedResult: 12.},
		{input: []float64{0.}, expectedResult: 0.},
		{input: []float64{-12.}, expectedResult: -12.},
		{input: []float64{-1., -1.}, expectedResult: -1.},
		{input: []float64{-1., 1.}, expectedResult: 0.},
		{input: []float64{1., 1.}, expectedResult: 1.},
		{input: []float64{1., 2., 3., 4., 5., 6.}, expectedResult: 3.5},
		{input: []float64{0., 1., -1., 2., -2.}, expectedResult: 0.0},
		{input: []float64{3.005, 3.005, 3.005, 3.005, 3.005, 3.005002, 3.004998}, expectedResult: 3.005},
	} {
		result, err := Mean(testData.input)
		if err != testData.expectedError {
			t.Errorf("%d) Wrong error value: input=%v, expected error=%v, obtained error=%v",
				testNumber, testData.input, testData.expectedError, err)
		}
		if !equals(result, testData.expectedResult, tolerance) {
			t.Errorf("%d) Wrong result: in=%v, out=%f, got=%f",
				testNumber, testData.input, testData.expectedResult, result)
		}
	}
}

func TestStandardDeviation(t *testing.T) {
	for testNumber, testData := range [...]struct {
		input          []float64
		expectedError  error
		expectedResult float64
	}{
		{input: nil, expectedError: ErrSampleTooSmall},
		{input: []float64{}, expectedError: ErrSampleTooSmall},
		{input: []float64{1.0, 1.0}, expectedResult: 0.0},
		{input: []float64{1.0, 2.0}, expectedResult: 0.707},
		{input: []float64{1.0, 2, 3, 4, 5, 6}, expectedResult: 1.870},
		{input: []float64{-2.0, -1, 0, 1, 2, 3}, expectedResult: 1.870},
	} {
		result, err := StandardDeviation(testData.input)

		if err != testData.expectedError {
			t.Errorf("%d) Wrong error value: input=%v, expected error=%v, obtained error=%v", testNumber, testData.input, testData.expectedError, err)
		}

		if !equals(result, testData.expectedResult, tolerance) {
			t.Errorf("%d) Wrong result value: input=%v, expected result=%f, obtained result=%f", testNumber, testData.input, testData.expectedResult, result)
		}
	}
}

func TestStandardError(t *testing.T) {
	for testNumber, testData := range [...]struct {
		input          []float64
		expectedError  error
		expectedResult float64
	}{
		{input: nil, expectedError: ErrSampleTooSmall},
		{input: []float64{}, expectedError: ErrSampleTooSmall},
		{input: []float64{1.0, 1.0}, expectedResult: 0.0},
		{input: []float64{1.0, 2.0}, expectedResult: 0.5},
		{input: []float64{1.0, 2, 3, 4, 5, 6}, expectedResult: 0.763},
		{input: []float64{-2.0, -1, 0, 1, 2, 3}, expectedResult: 0.763},
	} {
		result, err := StandardError(testData.input)

		if err != testData.expectedError {
			t.Errorf("%d) Wrong error value: input=%v, expected error=%v, obtained error=%v", testNumber, testData.input, testData.expectedError, err)
		}

		if !equals(result, testData.expectedResult, tolerance) {
			t.Errorf("%d) Wrong result value: input=%v, expected result=%f, obtained result=%f", testNumber, testData.input, testData.expectedResult, result)
		}
	}
}

func TestSumSamll(t *testing.T) {
	for i, f := range [...]struct {
		input    []float64
		expected float64
	}{
		{input: []float64{}, expected: 0.0},
		{input: []float64{3.1}, expected: 3.1},
		{input: []float64{3.1, 2.0}, expected: 5.1},
		{input: []float64{3.1, 2.0, 7.0}, expected: 12.1},
		{input: []float64{1.0, -1.0, 3.0}, expected: 3.0},
		{input: []float64{1.1, 2.0, 3.0, 4.0}, expected: 10.1},
	} {
		output := sum(f.input)
		if !equals(output, f.expected, tolerance) {
			t.Errorf("%d) input=%v, expected=%f, output=%f",
				i, f.input, f.expected, output)
		}
	}
}

func oneToN(n int) []float64 {
	s := make([]float64, n)
	for i := 0; i < n; i++ {
		s[i] = float64(i + 1)
	}
	return s
}

// this will tests the sum of big slices, using simple finite arithmetic
// progressions as the input data.
func TestSumBig(t *testing.T) {
	for i, n := range [...]int{
		10,
		100,
		1000,
		10000,
		100000,
	} {
		expected := float64(n*(1+n)) / 2
		input := oneToN(n)
		output := sum(input)
		if !equals(output, expected, tolerance) {
			t.Errorf("%d) n=%d, expected=%f, output=%f",
				i, n, expected, output)
		}
	}
}

var benchmarkResult float64 // avoid compiler optimization to elimitate tests

func benchmarkSum(n int, concurrent bool, b *testing.B) {
	b.StopTimer()
	input := oneToN(n)
	b.StartTimer()

	var r float64
	for i := 0; i < b.N; i++ {
		r = sum(input)
	}
	benchmarkResult = r
}

func BenchmarkSum4(b *testing.B) { benchmarkSum(1000, false, b) }
func BenchmarkSum5(b *testing.B) { benchmarkSum(10000, false, b) }
func BenchmarkSum6(b *testing.B) { benchmarkSum(100000, false, b) }
func BenchmarkSum7(b *testing.B) { benchmarkSum(1000000, false, b) }
func BenchmarkSum8(b *testing.B) { benchmarkSum(10000000, false, b) }

func BenchmarkSumConcurrent4(b *testing.B) { benchmarkSum(1000, true, b) }
func BenchmarkSumConcurrent5(b *testing.B) { benchmarkSum(10000, true, b) }
func BenchmarkSumConcurrent6(b *testing.B) { benchmarkSum(100000, true, b) }
func BenchmarkSumConcurrent7(b *testing.B) { benchmarkSum(1000000, true, b) }
func BenchmarkSumConcurrent8(b *testing.B) { benchmarkSum(10000000, true, b) }

func TestMeanConfidenceIntervalsErrors(t *testing.T) {
	for i, tt := range [...]struct {
		inData       []float64
		inConfidence float64
		expected     [2]float64
	}{
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.5, expected: [2]float64{1.0, 3.0}},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.90, expected: [2]float64{-4.3138, 8.3138}},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.95, expected: [2]float64{-10.71, 14.71}},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.99, expected: [2]float64{-61.66, 65.66}},

		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.50, expected: [2]float64{4.0928, 5.9072}},
		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.90, expected: [2]float64{2.3890, 7.6110}},
		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.95, expected: [2]float64{1.5996, 8.4004}},
		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.99, expected: [2]float64{-0.63884, 10.63884}},

		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.50, expected: [2]float64{-0.64395, -0.36088}},
		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.90, expected: [2]float64{-0.87162, -0.13321}},
		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.95, expected: [2]float64{-0.958031, -0.046796}},
		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.99, expected: [2]float64{-1.15696, 0.15213}},
	} {
		obtained, err := MeanConfidenceIntervals(tt.inData, tt.inConfidence)
		if err != nil {
			t.Errorf("%d) inData=%v, inConfidence=%f, expected=%v, obtained error=%q",
				i, tt.inData, tt.inConfidence, tt.expected, err)
		}

		if !pairEquals(obtained, tt.expected, tolerance) {
			t.Errorf("%d) inData=%v, inConfidence=%f, expected=%v, obtained=%v",
				i, tt.inData, tt.inConfidence, tt.expected, obtained)
		}
	}
}

func TestMeanConfidenceIntervalsObjErrors(t *testing.T) {
	for i, tt := range [...]struct {
		inData       []float64
		inConfidence float64
		expected     error
	}{
		{inData: []float64{1.0},
			inConfidence: 0.99, expected: ErrSampleTooSmall},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.0, expected: ErrInvalidConfidence},
		{inData: []float64{1.0, 3.0},
			inConfidence: 1.0, expected: ErrInvalidConfidence},
		{inData: []float64{1.0, 3.0},
			inConfidence: -1.0, expected: ErrInvalidConfidence},
		{inData: []float64{1.0, 3.0},
			inConfidence: 2.0, expected: ErrInvalidConfidence},
	} {
		_, err := MeanConfidenceIntervals(tt.inData, tt.inConfidence)
		if err != tt.expected {
			t.Errorf("%d) inData=%v, inConfidence=%f, expected=%q, obtained no error",
				i, tt.inData, tt.inConfidence, tt.expected)
		}

		if err != tt.expected {
			t.Errorf("%d) inData=%v, inConfidence=%f, expected=%q, obtained=%q",
				i, tt.inData, tt.inConfidence, tt.expected, err)
		}
	}
}

func TestMeanConfidenceIntervals(t *testing.T) {
	for i, tt := range [...]struct {
		inData       []float64
		inConfidence float64
		expected     [2]float64
	}{
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.5, expected: [2]float64{1.0, 3.0}},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.90, expected: [2]float64{-4.3138, 8.3138}},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.95, expected: [2]float64{-10.71, 14.71}},
		{inData: []float64{1.0, 3.0},
			inConfidence: 0.99, expected: [2]float64{-61.66, 65.66}},

		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.50, expected: [2]float64{4.0928, 5.9072}},
		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.90, expected: [2]float64{2.3890, 7.6110}},
		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.95, expected: [2]float64{1.5996, 8.4004}},
		{inData: []float64{2.0, 3, 5, 6, 9},
			inConfidence: 0.99, expected: [2]float64{-0.63884, 10.63884}},

		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.50, expected: [2]float64{-0.64395, -0.36088}},
		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.90, expected: [2]float64{-0.87162, -0.13321}},
		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.95, expected: [2]float64{-0.958031, -0.046796}},
		{inData: []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			inConfidence: 0.99, expected: [2]float64{-1.15696, 0.15213}},
	} {
		obtained, err := MeanConfidenceIntervals(tt.inData, tt.inConfidence)
		if err != nil {
			t.Errorf("%d) inData=%v, inConfidence=%f, expected=%v, obtained error=%q",
				i, tt.inData, tt.inConfidence, tt.expected, err)
		}

		if !pairEquals(obtained, tt.expected, tolerance) {
			t.Errorf("%d) inData=%v, inConfidence=%f, expected=%v, obtained=%v",
				i, tt.inData, tt.inConfidence, tt.expected, obtained)
		}
	}
}
