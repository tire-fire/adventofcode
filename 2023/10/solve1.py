#!/usr/bin/env python
import re
import numpy as np
from collections import deque

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]


def is_valid_pipe(grid, x, y, direction):
    if x < 0 or y < 0 or x >= len(grid) or y >= len(grid[0]):
        return False

    tile = grid[x][y]
    if direction == "N":
        return tile in {"|", "7", "F"}
    elif direction == "S":
        return tile in {"|", "L", "J"}
    elif direction == "E":
        return tile in {"-", "J", "7"}
    elif direction == "W":
        return tile in {"-", "L", "F"}


def find_max_distance(grid):
    # Find the starting position 'S'
    for i, row in enumerate(grid):
        for j, cell in enumerate(row):
            if cell == "S":
                start_x, start_y = i, j
                break

    # Directions: N, S, E, W
    directions = {"N": (-1, 0), "S": (1, 0), "E": (0, 1), "W": (0, -1)}
    visited = set()
    queue = deque([(start_x, start_y, 0)])  # (x, y, distance)
    max_distance = 0

    while queue:
        x, y, dist = queue.popleft()
        if (x, y) in visited:
            continue
        visited.add((x, y))

        # Update max distance
        max_distance = max(max_distance, dist)

        # Add adjacent tiles
        for d, (dx, dy) in directions.items():
            nx, ny = x + dx, y + dy
            if is_valid_pipe(grid, nx, ny, d) and (nx, ny) not in visited:
                queue.append((nx, ny, dist + 1))

    return max_distance


print(inp)
print(find_max_distance(inp))
