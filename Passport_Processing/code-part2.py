#!/usr/bin/env python3

import numpy as np
import re

required_fields = ["byr","iyr","eyr","hgt","hcl","ecl","pid"]
def get_num(x):
    try:
        dat = int(x)
        return dat
    except:
        return False

def val_byr(x):
        dat = get_num(x)
        if not dat:
            return False
        return dat <= 2002 and dat >= 1920

def val_iyr(x):
    dat = get_num(x)
    if not dat:
        return False
    return dat <= 2020 and dat >= 2010

def val_hgt(x):
    if "cm" in x:
        x = x[:-2]
        x = get_num(x)
        if not x:
            return False
        return x <= 193 and x >= 150
    if "in" in x:
        x = x[:-2]
        x = get_num(x)
        if not x:
            return False
        return x <= 76 and x >= 59
    return False

def val_hcl(x):
    if not("#" in x):
        return False
    x = x[1:]
    if len(x) > 6:
        return False
    pat = re.compile("^[0-9a-f]{6}$")
    return pat.search(x) is not None

def val_ecl(x):
    valids = ["amb","blu","brn","gry","grn","hzl","oth"]
    return x in valids

def val_pid(x):
    pat = re.compile("^\d{9}$")
    return pat.search(x) is not None
    
def val_eyr(x):
    dat = get_num(x)
    if not dat:
        return False
    return dat <= 2030 and dat >= 2020

def val_fields(lines):
    res = []
    for l in lines:
        res.append([x in l for x in required_fields])
    ret = list(map(all,res))
    #print(ret)
    #print(sum(ret))
    return ret

def get_dat(x):
    return x[4:]

def main():
    with open("data.txt") as f:
        lines = f.readlines()
        lines = "".join(lines)
    
    lines = lines.split("\n\n")
    fields = val_fields(lines)

    dat = []
    for l in lines:
        dat.append(l.split("\n"))

    res = []
    for d in dat:
        tmp = []
        for a in d:
            for itm in a.split(" "):
                tmp.append(itm)
        res.append(tmp)

    #["byr","iyr","eyr","hgt","hcl","ecl","pid"]
    fin_res = []
    for x in res:
        tmp = []
        for d in x:
            if "byr" in d:
                dat = get_dat(d)
                tmp.append(val_byr(dat))
            if "iyr" in d:
                dat = get_dat(d)
                tmp.append(val_iyr(dat))
            if "eyr" in d:
                dat = get_dat(d)
                tmp.append(val_eyr(dat))
            if "hgt" in d:
                dat = get_dat(d)
                tmp.append(val_hgt(dat))
            if "hcl" in d:
                dat = get_dat(d)
                tmp.append(val_hcl(dat))
            if "ecl" in d:
                dat = get_dat(d)
                tmp.append(val_ecl(dat))
            if "pid" in d:
                dat = get_dat(d)
                tmp.append(val_pid(dat))

        fin_res.append(tmp)

        fin_ret = list(map(all,fin_res))

    #print(fin_ret)
    #print(fields)

    ret = list(map(all,zip(fin_ret,fields)))
    print(ret)
    print(sum(ret))


if __name__ == "__main__":
    main()
