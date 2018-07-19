import math

def abs_diff(x, y):
    return abs(min(x, y) - max(x,y))

def scaleSteps(steps):
    raise NotImplementedError

def quantize(steps, x):
    stepIndex = None

    if abs(min(steps)) == abs(max(steps)):
        scale = max(steps)
    else:
        scale = max(scaleSteps(steps))

    scaled = x * scale

    distances = [abs_diff(scaled, step) for step in steps]

    return steps[distances.index(min(distances))]

def Q(steps, fs, func=quantize):
    return [func(steps, f) for f in fs]

# The output domain is the set of possible values that inputs can be mapped to.
output_domain = list(range(-5, 6, 1))

# This is an example input set, specifically crafted to match the output domain above.
# -1 should be mapped to -5, -0.8 should map to -4, etc.
fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

if __name__ == '__main__':
    o = Q(output_domain, fs)
    print(o)
