package main

import (
	"fmt"
	noise "github.com/therealfakemoot/genesis/noise"
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
func (d Domain) Quantize(fs []float64) []int {
	var ret []int

	steps := d.Steps()
	numSteps := float64(len(steps))

	quantize := func(x float64) int {
		if x >= 0.5 {
			return int(x*numSteps + 0)
		}
		return int(x*(numSteps-1) - 1)
	}

	// quantaSize := (d.Max - d.Min) / (math.Pow(2.0, stepFloat) - 1.0)

	for _, f := range fs {
		ret = append(ret, quantize(f))
	}

	return ret
}

func genFloats(x, y float64) []float64 {
	var fs []float64

	n := noise.NewWithSeed(8675309)

	for x := 0.0; x < 5.0; x++ {
		for y := 0.0; y < 5.0; y++ {
			fs = append(fs, n.Eval3(x, y, 0))
		}
	}

	return fs
}

func main() {
	d := Domain{
		Min:  -5.0,
		Max:  5.0,
		Step: 1.0,
	}

	fs := genFloats(5, 5)

	v := d.Quantize(fs)

	fmt.Printf("Output Domain: %+v\n", d)
	fmt.Printf("Input: %v\n", fs)
	fmt.Printf("Output: %v\n", v)
}
