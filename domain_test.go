package quantize

import (
	"reflect"
	"testing"
)

var (
	seed    int64 = 8675309
	edge          = -20.0
	mapMin        = 5.0
	mapMax        = 20.0
	mapStep       = 5.0
)

func buildTestDomains(edge float64) []Domain {
	var ret []Domain

	for x := edge; x < 0.0; x++ {
		ret = append(ret, Domain{Min: x, Max: -x})
	}
	return ret
}

func TestKnownValues(t *testing.T) {
	input := Domain{Min: -1, Max: 1}
	t.Run("Balanced Domains", func(t *testing.T) {
		t.Run("[-5,5]", func(t *testing.T) {
			d := Domain{Min: -5, Max: 5}
			fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
			quantized := QuantizeAll(fs, input, d)
			expected := []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}

			if !reflect.DeepEqual(quantized, expected) {
				t.Errorf("Expected: %v, got %v", expected, quantized)
			}
		})
		t.Run("[-100,100]", func(t *testing.T) {
			d := Domain{Min: -100, Max: 100}
			fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
			quantized := QuantizeAll(fs, input, d)
			expected := []float64{-100, -80, -60, -40, -20, 0, 20, 40, 60, 80, 100}

			if !reflect.DeepEqual(quantized, expected) {
				t.Errorf("Expected: %v, got %v", expected, quantized)
			}
		})
		t.Run("[-2500,2500]", func(t *testing.T) {
			t.Skip()
			d := Domain{Min: -2500, Max: 2500}
			fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
			quantized := QuantizeAll(fs, input, d)
			expected := []float64{-100, -80, -60, -40, -20, 0, 20, 40, 60, 80, 100}

			if !reflect.DeepEqual(quantized, expected) {
				t.Errorf("Expected: %v, got %v", expected, quantized)
			}
		})

	})

	d := Domain{Min: -5, Max: 5}
	fs := []float64{-1.0, -0.8, -0.6, -0.4, -0.2, 0, 0.2, 0.4, 0.6, 0.8, 1}
	quantized := QuantizeAll(fs, input, d)
	expected := []float64{-5, -4, -3, -2, -1, 0, 1, 2, 3, 4, 5}

	if !reflect.DeepEqual(quantized, expected) {
		t.Errorf("Expected: %v, got %v", expected, quantized)
	}
}
