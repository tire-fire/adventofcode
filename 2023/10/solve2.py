#!/usr/bin/env python
import re
import numpy as np
from collections import deque

# with open("example4", "r") as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]


def is_valid_pipe(grid, tile, x, y, direction, visited):
    if x < 0 or y < 0 or x >= len(grid) or y >= len(grid[0]) or visited[x][y]:
        return False

    valid_continuations = {
        "|": {
            "N": ["|", "7", "F"],
            "S": ["|", "L", "J"],
            "W": [],
            "E": [],
        },
        "-": {
            "N": [],
            "S": [],
            "W": ["-", "F", "L"],
            "E": ["-", "7", "J"],
        },
        "L": {
            "N": ["|", "7", "F"],
            "S": [],
            "W": [],
            "E": ["-", "J", "7"],
        },
        "J": {
            "N": ["|", "F", "7"],
            "S": [],
            "W": ["-", "L", "F"],
            "E": [],
        },
        "7": {
            "N": [],
            "S": ["|", "L", "J"],
            "W": ["-", "F", "L"],
            "E": [],
        },
        "F": {
            "N": [],
            "S": ["|", "L", "J"],
            "W": [],
            "E": ["-", "7", "J"],
        },
        ".": {"N": [], "S": [], "W": [], "E": []},
        "S": {
            "N": ["|", "F", "7"],
            "S": ["|", "L", "J"],
            "W": ["-", "F", "L"],
            "E": ["-", "7", "J"],
        },
    }

    next_tile = grid[x][y]

    # Check if the next tile is a valid continuation based on the direction and current tile
    return next_tile in valid_continuations[tile][direction]


def find_max_distance(grid):
    # Find the starting position 'S'
    for i, row in enumerate(grid):
        for j, cell in enumerate(row):
            if cell == "S":
                start_x, start_y = i, j
                break

    # Directions: N, S, E, W
    directions = {"N": (-1, 0), "S": (1, 0), "E": (0, 1), "W": (0, -1)}
    visited = [[False] * len(grid[0]) for _ in range(len(grid))]
    queue = deque([(start_x, start_y, 0)])  # (x, y, distance)
    max_distance = 0

    while queue:
        x, y, dist = queue.popleft()
        if (x, y) in visited:
            continue
        visited[x][y] = True

        # Update max distance
        max_distance = max(max_distance, dist)

        # Add adjacent tiles
        for d, (dx, dy) in directions.items():
            nx, ny = x + dx, y + dy
            if is_valid_pipe(grid, grid[x][y], nx, ny, d, visited):
                queue.append((nx, ny, dist + 1))

    return max_distance, visited


def count_enclosed_tiles(grid):
    max_distance, visited = find_max_distance(grid)

    # Count the enclosed tiles within the loop
    enclosed_count = 0
    for x in range(len(grid)):
        inside_network = False
        for y in range(len(grid[0])):
            if visited[x][y]:
                if (
                    grid[x][y] in "|JL"
                ):  # if it's crossed over a vertical an odd number of times
                    inside_network = not inside_network
            else:
                enclosed_count += inside_network  # then the .'s count as inside

    return enclosed_count


for line in inp:
    print(line)

print(count_enclosed_tiles(inp))
