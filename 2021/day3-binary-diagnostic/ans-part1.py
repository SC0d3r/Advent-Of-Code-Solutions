#!/usr/bin/env python3

import numpy as np

lines = []

with open("real-data.txt") as f:
    lines = f.readlines()

lines = list(map(lambda x:x.strip(), lines))

lines = list(map(lambda x: [int(y) for y in x], lines))

lines = np.array(lines)

gamma = ""
eps = ""

def leastCommon(xs):
    xs = list(xs)
    return min(set(xs), key=xs.count)

def mostCommon(xs):
    xs = list(xs)
    return max(set(xs), key=xs.count)

rows = lines.shape[0]
cols = lines.shape[1]

for i in range(cols):
    col = lines[:, i]
    gamma += str(mostCommon(col))
    eps += str(leastCommon(col))


print(eps, gamma)

print(int(eps, 2), int(gamma, 2))
        
print("final ans:", int(eps, 2) * int(gamma, 2))
