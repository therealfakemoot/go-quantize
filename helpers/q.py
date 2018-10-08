import math

def abs_diff(x, y):
    return abs(min(x, y) - max(x,y))

midpoint = lambda x: (x[0] + x[-1])/2.0

def balanceSteps(steps):
    m = midpoint(steps)
    shift = abs_diff(m, 0)

    if abs(min(steps)) < abs(max(steps)):
        shifted = [step - shift for step in steps]
    else:
        shifted = [step + shift for step in steps]

    return shifted

def Q(steps, fs):
    if abs(min(steps)) == abs(max(steps)):
        balanced = steps
        scale = max(steps)
    else:
        balanced = balanceSteps(steps)
        scale = max(balanced)

    def quantize(x):
        scaled = x * scale

        distances = [abs_diff(scaled, step) for step in balanced]

        return steps[distances.index(min(distances))]

    for f in fs:
        yield quantize(f)


if __name__ == '__main__':

    # The output domain is the set of possible values that inputs can be mapped to.
    output_domain = list(range(-5, 6, 1))

    # This is an example input set, specifically crafted to match the output domain above.
    # -1 should be mapped to -5, -0.8 should map to -4, etc.
    fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

    # o = Q(output_domain, fs)
    # print(o)
