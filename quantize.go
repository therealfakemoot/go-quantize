package quantize

import (
	"fmt"
)

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
