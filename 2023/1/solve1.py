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
    digits = re.findall(r"(\d|one|two|three|four|five|six|seven|eight|nine)", line)
    print(line)
    print(digits)
    first = digits[0] if digits[0].isdigit() else nums[digits[0]]
    second = digits[-1] if digits[-1].isdigit() else nums[digits[-1]]

    total += int(f"{first}{second}")
    print(total)
