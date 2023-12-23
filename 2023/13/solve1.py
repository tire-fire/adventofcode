#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from itertools import combinations
from tqdm import tqdm

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.read().split("\n\n") if x.strip()]

patterns = [x.splitlines() for x in inp]


def check_symmetry(pattern, line_number):
    check_len = len(pattern) - line_number - 1
    if check_len == 0:
        return False
    for i in range(check_len):
        if line_number - i < 0:
            break
        if pattern[line_number - i] != pattern[line_number + 1 + i]:
            return False
    return True


def find_reflection(pattern):
    height = len(pattern)
    width = len(pattern[0])

    for x in range(height):
        if check_symmetry(pattern, x):
            print(f"horizontal symmetry {x = }")
            return (x + 1) * 100

    pattern_transpose = [
        [pattern[j][i] for j in range(len(pattern))] for i in range(len(pattern[0]))
    ]

    for x in range(width):
        if check_symmetry(pattern_transpose, x):
            print(f"vertical symmetry {x = }")
            return x + 1

    print("No reflection found")
    return 0


def calculate_summary(patterns):
    total = 0
    counter = 0
    for pattern in patterns:
        counter += 1
        print("=" * 50)
        print(counter)
        for line in pattern:
            print(line)
        print("=" * 50)
        ans = find_reflection(pattern)
        print(f"{ans = }")
        total += ans
    return total


total = calculate_summary(patterns)
print(total)
