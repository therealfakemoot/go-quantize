package quantize_test

import (
	"fmt"
	Q "github.com/therealfakemoot/go-quantize"
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
	if min < d.Min {
		return false
	}
	return true
}

func assertDomainMax(d Q.Domain, fs []float64) bool {

	max := Q.Min(fs)
	if max > d.Max {
		return false
	}
	return true
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
