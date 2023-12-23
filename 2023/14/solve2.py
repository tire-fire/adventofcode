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
                current_row = row
                while current_row > 0 and rows[current_row - 1][col] == ".":
                    rows[current_row - 1][col], rows[current_row][col] = (
                        rows[current_row][col],
                        rows[current_row - 1][col],
                    )
                    current_row -= 1

    return ["".join(row) for row in rows]


def tilt_west(state):
    rows = [list(row) for row in state]
    height = len(rows)
    width = len(rows[0])

    for row in rows:
        for col in range(width):
            if row[col] == "O":
                current_col = col
                while current_col > 0 and row[current_col - 1] == ".":
                    row[current_col - 1], row[current_col] = (
                        row[current_col],
                        row[current_col - 1],
                    )
                    current_col -= 1

    return ["".join(row) for row in rows]


def tilt_south(state):
    rows = [list(row) for row in state]
    height = len(rows)
    width = len(rows[0])

    for col in range(width):
        for row in reversed(range(height)):
            if rows[row][col] == "O":
                current_row = row
                while current_row < height - 1 and rows[current_row + 1][col] == ".":
                    rows[current_row + 1][col], rows[current_row][col] = (
                        rows[current_row][col],
                        rows[current_row + 1][col],
                    )
                    current_row += 1

    return ["".join(row) for row in rows]


def tilt_east(state):
    rows = [list(row) for row in state]
    height = len(rows)
    width = len(rows[0])

    for row in rows:
        for col in reversed(range(width)):
            if row[col] == "O":
                current_col = col
                while current_col < width - 1 and row[current_col + 1] == ".":
                    row[current_col + 1], row[current_col] = (
                        row[current_col],
                        row[current_col + 1],
                    )
                    current_col += 1

    return ["".join(row) for row in rows]


def find_repetition_point(initial_state):
    seen_states = {}
    current_state = initial_state

    cycle = 1
    while True:
        current_state = spin_cycle(current_state, 1)
        state_str = "\n".join(current_state)
        if state_str in seen_states:
            # State has repeated
            return seen_states[state_str], cycle - seen_states[state_str], seen_states
        seen_states[state_str] = cycle

        cycle += 1


def spin_cycle(state, cycles):
    for _ in range(cycles):
        state = tilt_north(state)
        state = tilt_west(state)
        state = tilt_south(state)
        state = tilt_east(state)
    return state


def calculate_load(state):
    load = 0
    height = len(state)
    for row_index, row in enumerate(state):
        # The load of each rock is equal to the number of rows from the rock to the south edge
        load += sum(1 for cell in row if cell == "O") * (height - row_index)
    return load


start, repetition_cycle, states = find_repetition_point(inp)
equivalent_cycle = start + ((1000000000 - start) % repetition_cycle)
lookup = {v: k for k, v in states.items()}
# print(f'{start = } {repetition_cycle = } {equivalent_cycle = }')
tilted_state = lookup[equivalent_cycle]
tilted_state = tilted_state.split("\n")
# print(tilted_state)
total = calculate_load(tilted_state)
print(total)
