package quantize

import (
	"fmt"
)

// Domain describes the integer space to which float values must be mapped.
type Domain struct {
	Min  float64
	Max  float64
	Step float64
}

// Steps builds an array containing all steps in the desired output domain.
func (d Domain) Steps() []float64 {
	var steps []float64

	for i := d.Min; i <= d.Max; i += d.Step {
		steps = append(steps, i)
	}

	return steps
}

// Quantize normalizes a given set of arbitrary inputs into the provided output Domain.
func (d Domain) Quantize(fs []float64) []float64 {
	fmt.Println("Entering Quantize().")
	var ret = make([]float64, len(fs))

	for idx, f := range fs {
		ret[idx] = (f-(-1))*(d.Max-d.Min)/(1-(-1)) + d.Min
	}
	return ret
}
