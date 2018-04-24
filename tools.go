package quantize

import (
	noise "github.com/therealfakemoot/genesis/noise"
)

// Min returns the smallest member of an array of floats.
func Min(fs []float64) float64 {
	var min float64

	for _, v := range fs {
		if v < min {
			min = v
		}
	}

	return min
}

// Max returns the largest member of an array of floats.
func Max(fs []float64) float64 {
	var max float64

	for _, v := range fs {
		if v > max {
			max = v
		}
	}

	return max
}

// GenFloats will provide an x*y long array of floats, seeded with the given seed.
func GenFloats(x, y float64, seed int64) []float64 {
	var fs []float64

	n := noise.NewWithSeed(seed)

	for x := 0.0; x < 5.0; x++ {
		for y := 0.0; y < 5.0; y++ {
			fs = append(fs, n.Eval3(x, y, 0))
		}
	}

	return fs
}
