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

    # base case
    if len(group_sizes) == 0:
        return int("#" not in spring_row)

    total = 0
    # only look at the front group by removing the groups that have to be unique
    for position in range(
        len(spring_row)
        - sum(group_sizes[1:])
        + len(group_sizes[1:])
        - group_sizes[0]
        + 1
    ):
        # the first length can be operational springs, but still have to contain a group_size of broken springs followed by an operational spring separator
        possible = f'{"." * position}{"#" * group_sizes[0]}.'

        # check each point in the spring_row with the possible and see if it's valid
        if all(
            spring == possible_spring or spring == "?"
            for spring, possible_spring in zip(spring_row, possible)
        ):
            total += count_arrangements(
                spring_row[position + group_sizes[0] + 1 :], group_sizes[1:], memo
            )

    memo[(spring_row, group_sizes)] = total
    return total


memo = {}
total = 0
for row in inp:
    spring_row, group_sizes = row.split(" ")
    group_sizes = tuple(int(x) for x in group_sizes.split(","))
    unfold_row = "?".join([spring_row] * 5)
    unfold_group = group_sizes * 5
    total += count_arrangements(unfold_row, unfold_group, memo)

print(total)
