package quantize

import "math"

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
	var balanced []float64
	var scale float64

	steps := d.Steps()

	if math.Abs(Min(steps)) == math.Abs(Max(steps)) {
		balanced = steps
		scale = Max(steps)
	} else {
		balanced = balanceSteps(steps)
		scale = Max(balanced)
	}

	for _, val := range fs {
		var distances []float64
		scaled := val * scale
		for _, step := range balanced {
			distances = append(distances, absDiff(scaled, step))
		}
		minDist := Min(distances)

		for i, dist := range distances {
			if minDist == dist {
				ret = append(ret, steps[i])
			}
		}
	}

	return ret
}

func absDiff(x, y float64) float64 {
	return math.Abs(math.Min(x, y) - math.Max(x, y))
}

func midpoint(steps []float64) float64 {
	return (steps[0] + steps[len(steps)-1]) / 2
}

func balanceSteps(steps []float64) []float64 {
	var shifted []float64
	var shifter func(float64, float64) float64
	m := midpoint(steps)
	shift := absDiff(m, 0)

	if math.Abs(Min(steps)) < math.Abs(Max(steps)) {
		shifter = func(s, x float64) float64 { return x - s }
	} else {
		shifter = func(s, x float64) float64 { return x + s }
	}

	for _, s := range steps {
		shifted = append(shifted, shifter(s, shift))
	}

	return shifted
}
