#!/usr/bin/python3
import re

def get_num_regex(txt: str) -> int:
  nums = re.findall(r'\d', txt)
  return int(nums[0]  + nums[-1])

def main():
  with open("inp.txt") as f:
    res = sum([get_num_regex(l.strip()) for l in f.readlines()])
    print('res', res)

if __name__ == "__main__":
  main()