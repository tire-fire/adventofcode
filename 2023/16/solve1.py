#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from itertools import combinations
from tqdm import tqdm

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

UP, DOWN, LEFT, RIGHT = (-1, 0), (1, 0), (0, -1), (0, 1)


def get_new_direction(current_direction, tile):
    if tile == ".":
        return current_direction
    elif tile == "/":
        if current_direction == UP:
            return RIGHT
        elif current_direction == DOWN:
            return LEFT
        elif current_direction == RIGHT:
            return UP
        elif current_direction == LEFT:
            return DOWN
    elif tile == "\\":
        if current_direction == UP:
            return LEFT
        elif current_direction == DOWN:
            return RIGHT
        elif current_direction == RIGHT:
            return DOWN
        elif current_direction == LEFT:
            return UP
    elif tile in "|-":
        # Splitters, pass through if pointy end
        return current_direction


def simulate_beam(grid):
    energized_tiles = set()
    to_visit = [(0, 0, RIGHT)]  # Start at top-left, heading right
    visited = set()

    while to_visit:
        x, y, direction = to_visit.pop()
        while 0 <= x < len(grid) and 0 <= y < len(grid[0]):
            if (x, y, direction) in visited:
                break
            visited.add((x, y, direction))
            tile = grid[x][y]
            energized_tiles.add((x, y))

            if tile in "|-" and (x, y, direction) not in to_visit:
                # If it's a splitter add one to the queue and continue the other one
                if direction in [LEFT, RIGHT] and tile == "|":
                    to_visit.append((x, y, UP))
                    direction = DOWN
                elif direction in [UP, DOWN] and tile == "-":
                    to_visit.append((x, y, LEFT))
                    direction = RIGHT

            new_direction = get_new_direction(direction, tile)
            x, y = x + new_direction[0], y + new_direction[1]
            direction = new_direction

    return len(energized_tiles)


total = simulate_beam(inp)
print(total)
