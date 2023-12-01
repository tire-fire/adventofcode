#!/usr/bin/env python
import re

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

print(inp)

nums = {
    "zero": 0,
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9,
}

total = 0
for line in inp:
    numbers = []
    for i in range(len(line)):
        for word, digit in nums.items():
            if line[i:].startswith(word) or line[i:].startswith(str(digit)):
                numbers.append(digit)
    total += int(f"{numbers[0]}{numbers[-1]}")
    print(total)
