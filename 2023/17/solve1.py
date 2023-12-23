#!/usr/bin/env python
import re
import numpy as np
from itertools import combinations
from tqdm import tqdm
from heapq import heappush, heappop

# with open("example", "r") as f:
with open("input", "r") as f:
    inp = [[int(y) for y in x.strip()] for x in f.readlines() if x.strip()]

directions = {
    "S": [(0, 1, "R"), (1, 0, "D")],
    "L": [(1, 0, "D"), (-1, 0, "U")],
    "D": [(0, -1, "L"), (0, 1, "R")],
    "R": [(1, 0, "D"), (-1, 0, "U")],
    "U": [(0, -1, "L"), (0, 1, "R")],
}


def find_lowest_heat_path(grid):
    destination = (len(grid) - 1, len(grid[0]) - 1)
    heap = [(0, (0, 0), "S", "")]  # Heat cost, position, direction, path
    visited = set()

    while heap:
        heat_loss, (row, col), direction, path = heappop(heap)

        if (row, col) == destination:
            return heat_loss, path

        if (row, col, direction) in visited:
            continue

        visited.add((row, col, direction))

        for dr, dc, next_direction in directions[direction]:
            new_heat_loss = heat_loss
            for steps in range(1, 4):
                new_row, new_col = row + dr * steps, col + dc * steps

                if not (
                    0 <= new_row <= destination[0] and 0 <= new_col <= destination[1]
                ):
                    continue

                new_heat_loss = new_heat_loss + int(grid[new_row][new_col])
                heappush(
                    heap,
                    (
                        new_heat_loss,
                        (new_row, new_col),
                        next_direction,
                        path + next_direction,
                    ),
                )

    return None


result = find_lowest_heat_path(inp)
print(result)
