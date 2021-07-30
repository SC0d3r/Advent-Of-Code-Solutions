#!/usr/bin/env python3

import numpy as np

ps = []
required_fields = ["byr","iyr","eyr","hgt","hcl","ecl","pid"]
def main():
    with open("data.txt") as f:
        lines = f.readlines()
        lines = "".join(lines)
    
    lines = lines.split("\n\n")

    res = []
    for l in lines:
        res.append([x in l for x in required_fields])
    ret = list(map(all,res))
    print(ret)
    print(sum(ret))


if __name__ == "__main__":
    main()
