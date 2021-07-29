#!/usr/bin/env python3
import numpy as np

# You can find the question in here
# https://adventofcode.com/2020/day/3


def save(data):
    with open("tmp.txt", "w") as f:
        f.write(data)

def to_str(data):
    res = ""
    for x in data:
        res += "".join(x)
    return "".join(res)

def main():
    with open("data.txt") as f:
        path = [list(x.strip()) for x in f.readlines()]
        path = np.array(path)

    for _ in range(8):
        path = np.hstack((path,path))


    trees = 0
    ix, iy = [0,0]

    size_x, size_y = path.shape
    step_x = 1
    step_y = 2
    while True:
        ix += step_x
        iy += step_y 
        if ix >= size_y or iy >= size_x:
            break
        item = path[iy,ix]
        print(item)
        path[iy, ix] = "@"
        if item == "#": # aka tree
            trees += 1

    print("trees", trees)
        




if __name__ == "__main__":
    main()
