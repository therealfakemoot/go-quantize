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

func TestQuantizeMin(t *testing.T) {
	d := Domain{
		Min:  -5,
		Max:  5,
		Step: 1,
	}

	fs := GenFloats(20, 20, 8675309)

	quantized := d.Quantize(fs)

	t.Logf("Quantized Values: %v", quantized)
	min := MinInt(quantized)

	if min < int(d.Min) {
		t.Errorf("Domain Minimum (%d) exceeded: %d", int(d.Min), min)
	}

}

func TestQuantizeMax(t *testing.T) {
	d := Domain{
		Min:  -5,
		Max:  5,
		Step: 1,
	}

	fs := GenFloats(20, 20, 8675309)

	quantized := d.Quantize(fs)

	t.Logf("%v", quantized)
	max := MinInt(quantized)

	if max > int(d.Max) {
		t.Errorf("Domain Maximum (%d) exceeded: %d", int(d.Max), max)
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
