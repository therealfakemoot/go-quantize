package main

import (
	"fmt"
	noise "github.com/therealfakemoot/genesis/noise"
	// "math"
	// "math/rand"
)

// Domain describes the integer space to which float values must be mapped.
type Domain struct {
	Min  float64
	Max  float64
	Step float64
}

// func quantize(delta float64, i float64) float64 {
// return delta * math.Floor((i/delta)+.5)
// }

// Quantize normalizes a given set of arbitrary inputs into the provided output Domain.
func Quantize(d Domain, fs []float64) []int {
	var ret []int
	var steps []float64

	quantize := func(steps float64, x float64) int {
		if x >= 0.5 {
			return int(x*steps + 0)
		}
		return int(x*(steps-1) - 1)
	}

	for i := d.Min; i <= d.Max; i += d.Step {
		steps = append(steps, i)
	}

	stepFloat := float64(len(steps))
	// quantaSize := (d.Max - d.Min) / (math.Pow(2.0, stepFloat) - 1.0)

	for _, f := range fs {
		ret = append(ret, quantize(stepFloat, f))
	}

	// fmt.Printf("Steps: %v\n", steps)
	// fmt.Printf("Quanta size: %f\n", quantaSize)

	return ret
}

func main() {
	d := Domain{
		Min:  -5.0,
		Max:  5.0,
		Step: 1.0,
	}

	n := noise.NewWithSeed(8675309)

	var fs []float64
	for x := 0.0; x < 5.0; x++ {
		for y := 0.0; y < 5.0; y++ {
			fs = append(fs, n.Eval3(x, y, 0))
		}
	}

	// for i := 0; i < 20; i++ {
	// fs = append(fs, rand.Float64())
	// }

	v := Quantize(d, fs)

	fmt.Printf("Output Domain: %+v\n", d)
	fmt.Printf("%v\n", fs)
	fmt.Printf("%v\n", v)
}
