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


# calculate the OXR
oxr = lines.copy()
i = 0
while True:
    if len(oxr) == 1:
        break
    col = oxr[:, i]
    c = mostCommon(col)
    if list(col).count(1) == list(col).count(0):# is a tie choose 1
        c = 1
    oxr  = oxr[oxr[:, i] == c]
    i += 1

# calculate the CO2
co2  = lines.copy()
i = 0
while True:
    if len(co2) == 1:
        break
    col = co2[:, i]
    c = leastCommon(col)
    if list(col).count(1) == list(col).count(0):# is a tie choose 0
        c = 0
    co2  = co2[co2[:, i] == c]
    i += 1

def getNumber(res):
    return int("".join([str(x) for x in res[0,:]]), 2)

print("OXR:", getNumber(oxr))
print("co2", getNumber(co2))
print("Final ans:", getNumber(oxr) * getNumber(co2))
