package quantize_test

import (
	"fmt"
	Q "github.com/therealfakemoot/go-quantize"
	"reflect"
	"testing"
)

var (
	seed    int64 = 8675309
	edge          = -20.0
	mapMin        = 5.0
	mapMax        = 20.0
	mapStep       = 5.0
)

func buildTestDomains(edge float64) []Q.Domain {
	var ret []Q.Domain

	for x := edge; x < 0.0; x++ {
		ret = append(ret, Q.Domain{Min: x, Max: -x, Step: 1})
	}
	return ret
}

func assertDomainMin(d Q.Domain, fs []float64) bool {

	min := Q.Min(fs)
	return !(min < d.Min)
}

func assertDomainMax(d Q.Domain, fs []float64) bool {

	max := Q.Max(fs)
	return !(max > d.Max)
}

func TestDomainBounds(t *testing.T) {
	domains := buildTestDomains(edge)

	for _, d := range domains {
		t.Run(fmt.Sprintf("Domain:[%0.f,%0.f]", d.Min, d.Max), func(t *testing.T) {
			for i := mapMin; i <= mapMax; i += mapStep {
				fs := Q.Noise(i, i, seed)
				quantized := d.Quantize(fs)
				t.Run(fmt.Sprintf("Map(%0.f,%0.f:%d)", i, i, seed), func(t *testing.T) {
					t.Run("Min", func(t *testing.T) {
						if !assertDomainMin(d, quantized) {
							t.Fail()
						}
					})

					t.Run("Max", func(t *testing.T) {
						if !assertDomainMax(d, quantized) {
							t.Fail()
						}
					})

				})
			}
		})
	}

}

func TestKnownValues(t *testing.T) {
	t.Run("Balanced Domains", func(t *testing.T) {
		t.Run("[-5,5]", func(t *testing.T) {
			d := Q.Domain{Min: -5, Max: 5, Step: 1}
			fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
			quantized := d.Quantize(fs)
			expected := []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}

			if !reflect.DeepEqual(quantized, expected) {
				t.Errorf("Expected: %v, got %v", expected, quantized)
			}
		})
		t.Run("[-100,100]", func(t *testing.T) {
			d := Q.Domain{Min: -100, Max: 100, Step: 1}
			fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
			quantized := d.Quantize(fs)
			expected := []float64{-100, -80, -60, -40, -20, 0, 20, 40, 60, 80, 100}

			if !reflect.DeepEqual(quantized, expected) {
				t.Errorf("Expected: %v, got %v", expected, quantized)
			}
		})
		t.Run("[-2500,2500]", func(t *testing.T) {
			t.Skip()
			d := Q.Domain{Min: -2500, Max: 2500, Step: 1}
			fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
			quantized := d.Quantize(fs)
			expected := []float64{-100, -80, -60, -40, -20, 0, 20, 40, 60, 80, 100}

			if !reflect.DeepEqual(quantized, expected) {
				t.Errorf("Expected: %v, got %v", expected, quantized)
			}
		})

	})

	d := Q.Domain{Min: -5, Max: 5, Step: 1}
	fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
	quantized := d.Quantize(fs)
	expected := []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}

	if !reflect.DeepEqual(quantized, expected) {
		t.Errorf("Expected: %v, got %v", expected, quantized)
	}
}

func TestDomainSteps(t *testing.T) {
	d := Q.Domain{
		Min:  -5,
		Max:  5,
		Step: 1,
	}

	steps := d.Steps()

	if len(steps) != 11 {
		t.Errorf("Expected 11 steps, got: %d", len(steps))
	}
}
