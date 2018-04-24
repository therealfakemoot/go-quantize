package quantize

import (
	"fmt"
	noise "github.com/therealfakemoot/genesis/noise"
)

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

func main() {
	d := Domain{
		Min:  -5.0,
		Max:  5.0,
		Step: 1.0,
	}

	fs := GenFloats(5, 5, 8675309)

	v := d.Quantize(fs)

	fmt.Printf("Output Domain: %+v\n", d)
	fmt.Printf("Input: %v\n", fs)
	fmt.Printf("Output: %v\n", v)
}
