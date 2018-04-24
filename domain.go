package quantize

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
func (d Domain) Quantize(fs []float64) []float64 {
	var ret []float64

	steps := d.Steps()
	numSteps := float64(len(steps))

	quantize := func(x float64) float64 {
		if x >= 0.5 {
			return x*numSteps + 0
		}
		return x*(numSteps-1) - 1
	}

	// quantaSize := (d.Max - d.Min) / (math.Pow(2.0, stepFloat) - 1.0)

	for _, f := range fs {
		ret = append(ret, quantize(f))
	}

	return ret
}