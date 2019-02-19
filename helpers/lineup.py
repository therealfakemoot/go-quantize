import numpy as np
import matplotlib.pyplot as plt
import matplotlib.ticker as ticker

points = lambda x,y,z: [i/10 for i in range(x*10,(y*10)+1,int(z*10))]

def setupBase(ax):
    ax.spines['right'].set_color('none')
    ax.spines['left'].set_color('none')
    ax.yaxis.set_major_locator(ticker.NullLocator())
    ax.spines['top'].set_color('none')
    ax.xaxis.set_ticks_position('bottom')
    ax.tick_params(which='major', width=1.00)
    ax.tick_params(which='major', length=5)
    ax.tick_params(which='minor', width=0.75)
    ax.tick_params(which='minor', length=2.5)
    ax.set_ylim(0, 1)
    ax.patch.set_alpha(0.0)

def setupDomain(ax):
    ax.spines['right'].set_color('none')
    ax.spines['left'].set_color('none')
    ax.yaxis.set_major_locator(ticker.NullLocator())
    ax.spines['bottom'].set_color('none')
    ax.xaxis.set_ticks_position('bottom')
    ax.tick_params(which='major', width=1.00)
    ax.tick_params(which='major', length=5)
    ax.tick_params(which='minor', width=0.75)
    ax.tick_params(which='minor', length=2.5)
    ax.set_ylim(0, 1)
    ax.patch.set_alpha(0.0)

plt.figure(figsize=(8,6))

p = points(-1, 1, 0.1)
y = [0 for i in p]

base = plt.subplot(8, 1, 1)
setupBase(base)
base.set_xlim(min(p), max(p))

# domain = plt.subplot(9, 1, 1)
# setupDomain(domain)
# domain.set_xlim(-5, 5)

fs = [1.9052595476929043e-65, -0.24547524871149487, -0.14958647968356706, 0.2524271844660192, 0.07543249830197025, -0.2358464181549405, -2.0927125031917226e-65, -0.31351632106756133, 0.43597426984697757, 0.06148867313915843]

plt.plot(fs, y[:len(fs)], "r|")
plt.savefig('/var/www/idle.ndumas.com/line.svg')
