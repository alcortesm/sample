package sample

import (
	"fmt"
	"math"
	"testing"
)

const tolerance = 1e-3

func equals(a, b, tolerance float64) bool {
	return math.Abs(a-b) < tolerance
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
	t.Parallel()
	for _, test := range []struct {
		input []float64
		want  float64
	}{
		{
			input: []float64{12},
			want:  12,
		}, {
			input: []float64{0},
			want:  0,
		}, {
			input: []float64{-12},
			want:  -12,
		}, {
			input: []float64{-1, -1},
			want:  -1,
		}, {
			input: []float64{-1, 1},
			want:  0,
		}, {
			input: []float64{1, 1},
			want:  1,
		}, {
			input: []float64{1, 2, 3, 4, 5, 6},
			want:  3.5,
		}, {
			input: []float64{0, 1, -1, 2, -2},
			want:  0,
		}, {
			input: []float64{3.005, 3.005, 3.005, 3.005, 3.005, 3.005002, 3.004998},
			want:  3.005,
		},
	} {
		test := test
		description := fmt.Sprint(test.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := Mean(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if !equals(got, test.want, tolerance) {
				t.Errorf("wrong result: want %f, got %f", test.want, got)
			}
		})
	}
}

func TestMeanError(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		description string
		input       []float64
		want        error
	}{
		{
			description: "nil input",
			input:       nil,
			want:        ErrSampleTooSmall,
		}, {
			description: "empty input",
			input:       []float64{},
			want:        ErrSampleTooSmall,
		},
	} {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			_, err := Mean(test.input)
			if err == nil {
				t.Fatalf("unexpected success")
			}
			if err != test.want {
				t.Errorf("want %q, got %q", test.want, err)
			}
		})
	}
}

func TestStandardDeviation(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		input []float64
		want  float64
	}{
		{input: []float64{1, 1}, want: 0},
		{input: []float64{1, 2}, want: 0.707},
		{input: []float64{1, 2, 3, 4, 5, 6}, want: 1.870},
		{input: []float64{-2, -1, 0, 1, 2, 3}, want: 1.870},
	} {
		test := test
		description := fmt.Sprint(test.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := StandardDeviation(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if !equals(got, test.want, tolerance) {
				t.Errorf("want %f, got %f", test.want, got)
			}
		})
	}
}

func TestStandardDeviationError(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		description string
		input       []float64
		want        error
	}{
		{
			description: "nil input",
			input:       nil,
			want:        ErrSampleTooSmall,
		}, {
			description: "empty input",
			input:       []float64{},
			want:        ErrSampleTooSmall,
		},
	} {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			_, err := StandardDeviation(test.input)
			if err == nil {
				t.Fatal("unexpected success")
			}
			if err != test.want {
				t.Errorf("want %f, got %f", test.want, err)
			}
		})
	}
}

func TestStandardError(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		input []float64
		want  float64
	}{
		{
			input: []float64{1.0, 1.0},
			want:  0.0,
		}, {
			input: []float64{1.0, 2.0},
			want:  0.5,
		}, {
			input: []float64{1.0, 2, 3, 4, 5, 6},
			want:  0.763,
		}, {
			input: []float64{-2.0, -1, 0, 1, 2, 3},
			want:  0.763,
		},
	} {
		test := test
		description := fmt.Sprint(test.input)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := StandardError(test.input)
			if err != nil {
				t.Fatal(err)
			}
			if !equals(got, test.want, tolerance) {
				t.Errorf("want %f, got %f", test.want, got)
			}
		})
	}
}

func TestStandardErrorError(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		description string
		input       []float64
		want        error
	}{
		{
			description: "nil input",
			input:       nil,
			want:        ErrSampleTooSmall,
		}, {
			description: "empty input",
			input:       []float64{},
			want:        ErrSampleTooSmall,
		},
	} {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			_, err := StandardError(test.input)
			if err == nil {
				t.Errorf("unexpected success")
			}
			if err != test.want {
				t.Errorf("want %q, got %q", test.want, err)
			}
		})
	}
}

func TestMeanConfidenceIntervals(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		data       []float64
		confidence float64
		want       [2]float64
	}{
		{
			data:       []float64{1.0, 3.0},
			confidence: 0.5,
			want:       [2]float64{1.0, 3.0},
		},
		{
			data:       []float64{1.0, 3.0},
			confidence: 0.90,
			want:       [2]float64{-4.3138, 8.3138},
		}, {
			data:       []float64{1.0, 3.0},
			confidence: 0.95,
			want:       [2]float64{-10.71, 14.71},
		}, {
			data:       []float64{1.0, 3.0},
			confidence: 0.99,
			want:       [2]float64{-61.66, 65.66},
		}, {
			data:       []float64{2.0, 3, 5, 6, 9},
			confidence: 0.50,
			want:       [2]float64{4.0928, 5.9072},
		}, {
			data:       []float64{2.0, 3, 5, 6, 9},
			confidence: 0.90,
			want:       [2]float64{2.3890, 7.6110},
		}, {
			data:       []float64{2.0, 3, 5, 6, 9},
			confidence: 0.95,
			want:       [2]float64{1.5996, 8.4004},
		}, {
			data:       []float64{2.0, 3, 5, 6, 9},
			confidence: 0.99,
			want:       [2]float64{-0.63884, 10.63884},
		}, {
			data:       []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			confidence: 0.50,
			want:       [2]float64{-0.64395, -0.36088},
		}, {
			data:       []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			confidence: 0.90,
			want:       [2]float64{-0.87162, -0.13321},
		}, {
			data:       []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			confidence: 0.95,
			want:       [2]float64{-0.958031, -0.046796},
		}, {
			data:       []float64{-1.164837, -0.603101, -1.122721, -0.716435, 0.049454, 0.097798, 0.396846, -1.558289, -0.231544, -0.171306},
			confidence: 0.99,
			want:       [2]float64{-1.15696, 0.15213},
		},
	} {
		test := test
		description := fmt.Sprintf("%v, %v", test.data, test.confidence)
		t.Run(description, func(t *testing.T) {
			t.Parallel()
			got, err := MeanConfidenceIntervals(test.data, test.confidence)
			if err != nil {
				t.Fatal(err)
			}
			if !pairEquals(got, test.want, tolerance) {
				t.Errorf("want %f, got %f", test.want, got)
			}
		})
	}
}

func TestMeanConfidenceIntervalsErrors(t *testing.T) {
	t.Parallel()
	for _, test := range []struct {
		data       []float64
		confidence float64
		want       error
	}{
		{
			data:       []float64{1.0},
			confidence: 0.99,
			want:       ErrSampleTooSmall,
		}, {
			data:       []float64{1.0, 3.0},
			confidence: 0.0,
			want:       ErrInvalidConfidence,
		}, {
			data:       []float64{1.0, 3.0},
			confidence: 1.0,
			want:       ErrInvalidConfidence,
		}, {
			data:       []float64{1.0, 3.0},
			confidence: -1.0,
			want:       ErrInvalidConfidence,
		}, {
			data:       []float64{1.0, 3.0},
			confidence: 2.0,
			want:       ErrInvalidConfidence,
		},
	} {
		test := test
		description := fmt.Sprintf("%v, %v", test.data, test.confidence)
		t.Run(description, func(t *testing.T) {
			_, err := MeanConfidenceIntervals(test.data, test.confidence)
			if err == nil {
				t.Fatal("unexpected success")
			}
			if err != test.want {
				t.Errorf("want %q, got %q", test.want, err)
			}
		})
	}
}
