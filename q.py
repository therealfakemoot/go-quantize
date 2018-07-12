import math

# sign = lambda x: (1, -1)[x < 0]

def quantize(steps, x):
    distances = []
    stepIndex = None

    # scale = (len(steps)/(len(steps)-1))

    for step in steps:
        # distances.append(step - x)
        distances.append(abs(step) - abs(x))

    minDistance = min(distances)
    print("Value:{0}\nDistances:{1}\nMinDistance:{2}".format(x, distances, minDistance))

    return steps[distances.index(minDistance)]

def Q(steps, fs, func=quantize):
    return [func(steps, f) for f in fs]

# The output domain is the set of possible values that inputs can be mapped to.
output_domain = list(range(-5, 6, 1))

# This is an example input set, specifically crafted to match the output domain above.
# -1 should be mapped to -5, -0.8 should map to -4, etc.
fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]
