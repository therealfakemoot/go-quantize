package main

import (
	"testing"
)

func MinFloat(ints []float64) float64 {
	var min float64

	for _, v := range ints {
		if v < min {
			min = v
		}
	}

	return min
}

func MaxFloat(ints []float64) float64 {
	var max float64

	for _, v := range ints {
		if v > max {
			max = v
		}
	}

	return max
}

func prepTests(edge float64) []Domain {
	var ret []Domain

	for x := edge; x < 0.0; x++ {
		ret = append(ret, Domain{Min: x, Max: -x, Step: 1})
	}
	return ret
}

func TestDomainMin(t *testing.T) {
	domains := prepTests(-20)

	for _, d := range domains {

		for i := 5.0; i <= 20.0; i += 5 {
			fs := GenFloats(i, i, 8675309)

			quantized := d.Quantize(fs)

			t.Logf("Domain: %+v", d)
			t.Logf("GenFloats(%f, %f)", i, i)
			t.Logf("Quantized Values: %v", quantized)

			min := MinFloat(quantized)

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
			fs := GenFloats(i, i, 8675309)

			quantized := d.Quantize(fs)

			t.Logf("Domain: %+v", d)
			t.Logf("GenFloats(%f, %f)", i, i)
			t.Logf("Quantized Values: %v", quantized)
			max := MaxFloat(quantized)

			if max > d.Max {
				t.Errorf("Domain Maximum (%f) exceeded: %f", d.Max, max)
			}
		}
	}
}

func TestDomainSteps(t *testing.T) {
	d := Domain{
		Min:  -5,
		Max:  5,
		Step: 1,
	}

	steps := d.Steps()

	if len(steps) != 11 {
		t.Errorf("Expected 11 steps, got: %d", len(steps))
	}
}
