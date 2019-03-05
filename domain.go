package quantize

import (
	"fmt"
)

// Domain describes the integer space to which float values must be mapped.
type Domain struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
}

func (d Domain) String() string {
	return fmt.Sprintf("%#v", d)
}

func Quantize(f float64, input, output Domain) float64 {
	return (f-(input.Min))*(output.Max-output.Min)/(input.Max-input.Min) + output.Min

}

// Quantize normalizes a given set of arbitrary inputs into the provided output Domain.
func QuantizeAll(fs []float64, input, output Domain) []float64 {
	var ret = make([]float64, len(fs))

	for idx, f := range fs {
		ret[idx] = Quantize(f, input, output)
	}
	return ret
}
