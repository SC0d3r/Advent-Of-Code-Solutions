#!/usr/bin/env python3

lines = []

with open("real-data.txt") as f:
    lines = f.readlines()
lines = list(map(lambda x:x.strip(), lines))

print(lines)

def down(cLoc, y):
    cLoc['aim'] += y

def up(cLoc, y):
    cLoc['aim'] -= y
    
def forward(cLoc, x):
    cLoc['x'] += x
    cLoc['y'] += cLoc["aim"] * x

actions = {"forward":forward,"up":up,"down":down}


def getXY(data):
    return int(data.split()[1])

def getAction(data):
    return data.split()[0]


currentLoc = {"x":0, "y":0, "aim":0}  

for l in lines:
    xy = getXY(l)
    action = getAction(l)
    actions[action](currentLoc, xy)


print(currentLoc, currentLoc["x"] * currentLoc["y"])

