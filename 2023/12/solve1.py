#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from itertools import combinations
from tqdm import tqdm

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]


def count_arrangements(spring_row, group_sizes, memo):
    if (spring_row, group_sizes) in memo:
        return memo[(spring_row, group_sizes)]

    def is_valid(combination):
        count, groups = 0, []
        for char in combination:
            if char == "#":
                count += 1
            elif count > 0:
                groups.append(count)
                count = 0
        if count > 0:
            groups.append(count)
        return tuple(groups) == group_sizes

    def generate_combinations(index, spring_row):
        if index >= len(spring_row):
            if is_valid(spring_row):
                return 1
            return 0

        count = 0
        new_row = list(spring_row)
        if new_row[index] == "?":
            for replacement in [".", "#"]:
                new_row[index] = replacement
                count += generate_combinations(index + 1, "".join(new_row))
            new_row[index] = "?"
        else:
            count += generate_combinations(index + 1, spring_row)

        return count

    total = generate_combinations(0, spring_row)
    memo[(spring_row, group_sizes)] = total
    return total


memo = {}
total = 0
for row in inp:
    spring_row, group_sizes = row.split(" ")
    group_sizes = tuple(int(x) for x in group_sizes.split(","))
    total += count_arrangements(spring_row, group_sizes, memo)

print(total)
