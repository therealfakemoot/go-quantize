go-quantize converts an arbitrary sequence of floating point inputs into a well defined output space.

# Use Case

One use case for a numeric quantization function is to normalize inputs into a more reasonable output space. If a given function or set of measurements produces values too small or densely packed for human "consumption"/analysis, quantizing them into an integer space would allow a human to "eyeball" the distinctions between these values more sensible.

# Caveats
Quantization is a *lossy* conversion. The quantization function is in effect a hash function and depending on your inputs and output parameters, multiple input values may be mapped to the same output value. Be aware of this and do not rely on perfect fidelity of input->output mappings.

# Example

For the purposes of this example, assume that `GenFloats` is a function that produces values in the interval [-1,1].


```go
import "fmt"

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
```

The return value of `Quantize()` is an array of float64 values, normalized to the given Domain: in the case of the example, there will be no values smaller than -5, no values greater than 5.0, and all values will be integer increments between the two end points: -5, -4, -3, -2, -1, 0, 1, 2, 3, 4, and 5.
