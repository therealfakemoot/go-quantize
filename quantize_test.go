package main

import (
	"testing"
)

func MinInt(ints []int) int {
	var min int

	for _, v := range ints {
		if v < min {
			min = v
		}
	}

	return min
}

func MaxInt(ints []int) int {
	var max int

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

func TestQuantizeMin(t *testing.T) {
	domains := prepTests(-200)

	for _, d := range domains {

		for i := 5.0; i <= 200.0; i += 5 {
			fs := GenFloats(i, i, 8675309)

			quantized := d.Quantize(fs)

			t.Logf("Domain: %+v", d)
			t.Logf("GenFloats(%f, %f)", i, i)
			t.Logf("Quantized Values: %v", quantized)

			min := MinInt(quantized)

			if min < int(d.Min) {
				t.Errorf("Domain Minimum (%d) exceeded: %d\n", int(d.Min), min)
			}
		}
	}

}

func TestQuantizeMax(t *testing.T) {
	domains := prepTests(-200)

	for _, d := range domains {

		for i := 5.0; i <= 200.0; i += 5 {
			fs := GenFloats(i, i, 8675309)

			quantized := d.Quantize(fs)

			t.Logf("Domain: %+v", d)
			t.Logf("GenFloats(%f, %f)", i, i)
			t.Logf("Quantized Values: %v", quantized)
			max := MinInt(quantized)

			if max > int(d.Max) {
				t.Errorf("Domain Maximum (%d) exceeded: %d", int(d.Max), max)
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
