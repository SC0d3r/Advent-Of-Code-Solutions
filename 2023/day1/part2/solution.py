#!/usr/bin/python3
import re

word_number = {
  "one": "1",
  "two": "2",
  "three": "3",
  "four": "4",
  "five": "5",
  "six": "6",
  "seven": "7",
  "eight": "8",
  "nine": "9"
}

def get_num_regex(txt: str) -> int:
  nums = re.findall(r'\d|one|two|three|four|five|six|seven|eight|nine', txt)
  
  # replace the words with numbers
  for i, x in enumerate(nums):
    x = x.lower()
    if x in word_number:
      nums[i] = word_number[x]

  return int(nums[0]  + nums[-1])

def main():
  with open("inp.txt") as f:
    res = sum([get_num_regex(l.strip()) for l in f.readlines()])
    print('res', res)

if __name__ == "__main__":
  main()