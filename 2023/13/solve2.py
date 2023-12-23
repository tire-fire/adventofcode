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


def find_reflection(pattern):
    height = len(pattern)
    width = len(pattern[0])

    original_horizontal = [
        check_symmetry(pattern, mid_row) for mid_row in range(height)
    ]
    original_vertical = [
        check_symmetry(pattern, mid_col, is_vertical=True) for mid_col in range(width)
    ]

    for r in range(height):
        for c in range(width):
            flipped_pattern = flip_character(pattern, r, c)

            for mid_row in range(height):
                if (
                    check_symmetry(flipped_pattern, mid_row)
                    and not original_horizontal[mid_row]
                ):
                    return (mid_row + 1) * 100

            for mid_col in range(width):
                if (
                    check_symmetry(flipped_pattern, mid_col, is_vertical=True)
                    and not original_vertical[mid_col]
                ):
                    return mid_col + 1

    return 0


def flip_character(pattern, row, col):
    flipped_pattern = [list(line) for line in pattern]
    flipped_pattern[row][col] = "." if pattern[row][col] == "#" else "#"
    return ["".join(line) for line in flipped_pattern]


def check_symmetry(pattern, line_number, is_vertical=False):
    if is_vertical:
        pattern = ["".join(row[i] for row in pattern) for i in range(len(pattern[0]))]
    check_len = len(pattern) - line_number - 1
    for i in range(check_len):
        if line_number - i < 0 or line_number + i >= len(pattern):
            break
        if pattern[line_number - i] != pattern[line_number + 1 + i]:
            return False
    return True


def calculate_summary(patterns):
    total = 0
    for pattern in patterns:
        total += find_reflection(pattern)
    return total


total = calculate_summary(patterns)
print(total)
