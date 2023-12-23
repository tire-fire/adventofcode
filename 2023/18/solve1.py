#!/usr/bin/env python
import re
import numpy as np
from itertools import combinations
from tqdm import tqdm
from heapq import heappush, heappop

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

dig_plan = []
for line in inp:
    x, y, z = line.split(" ")
    dig_plan.append((x, int(y), z[1:-1]))

digger_pos = [0, 0]
grid = set()
directions = {"U": (0, 1), "D": (0, -1), "L": (-1, 0), "R": (1, 0)}

for direction, distance, _ in dig_plan:
    dx, dy = directions[direction]
    for _ in range(distance):
        digger_pos[0] += dx
        digger_pos[1] += dy
        grid.add(tuple(digger_pos))


def adjust_grid(grid):
    min_x = min(grid, key=lambda p: p[0])[0]
    min_y = min(grid, key=lambda p: p[1])[1]
    adjusted_grid = {(x - min_x + 1, y - min_y + 1) for x, y in grid}
    max_x = max(adjusted_grid, key=lambda p: p[0])[0] + 1
    max_y = max(adjusted_grid, key=lambda p: p[1])[1] + 1
    return adjusted_grid, min_x, min_y, max_x, max_y


def create_matrix(grid, max_x, max_y):
    width, height = max_x + 1, max_y + 1
    matrix = [["." for _ in range(width)] for _ in range(height)]
    for x, y in grid:
        matrix[y][x] = "#"
    return matrix


def flood_fill(grid, start, matrix, fill_char):
    to_fill = [start]
    while to_fill:
        x, y = to_fill.pop()
        if not (0 <= y < len(matrix) and 0 <= x < len(matrix[0])):
            continue
        if matrix[y][x] != ".":
            continue
        matrix[y][x] = fill_char
        to_fill.extend([(x + 1, y), (x - 1, y), (x, y + 1), (x, y - 1)])


# Adjust the grid, create the matrix, and perform the flood fill
adjusted_grid, min_x, min_y, max_x, max_y = adjust_grid(grid)
matrix = create_matrix(adjusted_grid, max_x, max_y)
exterior_start_point = (0, 0)  # point outside the trench
flood_fill(adjusted_grid, exterior_start_point, matrix, "o")

interior_count = sum(row.count(".") for row in matrix)
grid_drawing = "\n".join("".join(row) for row in matrix)

# Output
print(grid_drawing)
print(len(grid) + interior_count)
