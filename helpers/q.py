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
    low, high = min(steps), max(steps)
    for f in fs:
        yield (f-(-1))*(high-low)/(1-(-1)) + low


if __name__ == '__main__':

    # The output domain is the set of possible values that inputs can be mapped to.
    output_domain = list(range(-30, 16, 1))

    # This is an example input set, specifically crafted to match the output domain above.
    # -1 should be mapped to -5, -0.8 should map to -4, etc.
    fs = [-1, -.8, -.6, -.4, -.2, 0, .2, .4, .6, .8, 1]

    # fs = [1.9052595476929043e-65,-0.24547524871149487,-0.14958647968356706,0.2524271844660192,0.07543249830197025,-0.2358464181549405,-2.0927125031917226e-65,-0.31351632106756133,0.43597426984697757,0.06148867313915843]

    o = Q(output_domain, fs)
    # print(o)
