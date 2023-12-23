#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from itertools import combinations
from tqdm import tqdm

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]


def tilt_north(state):
    height = len(state)
    width = len(state[0])
    rows = [list(row) for row in state]

    for col in range(width):
        for row in range(height):
            if rows[row][col] == "O":
                # Move the rock up until it hits a cube-shaped rock or reaches the top
                current_row = row
                while current_row > 0 and rows[current_row - 1][col] == ".":
                    rows[current_row - 1][col], rows[current_row][col] = (
                        rows[current_row][col],
                        rows[current_row - 1][col],
                    )
                    current_row -= 1

    return ["".join(row) for row in rows]


tilted_state = tilt_north(inp)
for line in tilted_state:
    print(line)


def calculate_load(state):
    load = 0
    height = len(state)
    for row_index, row in enumerate(state):
        # The load of each rock is equal to the number of rows from the rock to the south edge
        load += sum(1 for cell in row if cell == "O") * (height - row_index)
    return load


total = calculate_load(tilted_state)
print(total)
