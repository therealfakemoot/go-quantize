package quantize_test

import (
	Q "github.com/therealfakemoot/go-quantize"
	"testing"
)

func prepTests(edge float64) []Q.Domain {
	var ret []Q.Domain

	for x := edge; x < 0.0; x++ {
		ret = append(ret, Q.Domain{Min: x, Max: -x, Step: 1})
	}
	return ret
}

func TestDomainMin(t *testing.T) {
	domains := prepTests(-20)

	for _, d := range domains {

		for i := 5.0; i <= 20.0; i += 5 {
			fs := Q.GenFloats(i, i, 8675309)

			quantized := d.Quantize(fs)

			t.Logf("Domain: %+v", d)
			t.Logf("GenFloats(%f, %f)", i, i)
			t.Logf("Quantized Values: %v", quantized)

			min := Q.Min(quantized)

			if min < d.Min {
				t.Errorf("Domain Minimum (%f) exceeded: %f\n", d.Min, min)
			}
		}
	}

}

func TestDomainMax(t *testing.T) {
	domains := prepTests(-20)

	for _, d := range domains {

		for i := 5.0; i <= 20.0; i += 5 {
			fs := Q.GenFloats(i, i, 8675309)

			quantized := d.Quantize(fs)

			t.Logf("Domain: %+v", d)
			t.Logf("GenFloats(%f, %f)", i, i)
			t.Logf("Quantized Values: %v", quantized)
			max := Q.Max(quantized)

			if max > d.Max {
				t.Errorf("Domain Maximum (%f) exceeded: %f", d.Max, max)
			}
		}
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
