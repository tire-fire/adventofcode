#!/usr/bin/env python
import re
import numpy as np
from itertools import combinations
from tqdm import tqdm
from heapq import heappush, heappop

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

directions = {"R": (1, 0), "D": (0, -1), "L": (-1, 0), "U": (0, 1)}
dig_plan = []
for line in inp:
    x, y, z = line.split(" ")
    distance = int(z[2:-2], 16)
    direction = list(directions.keys())[int(z[-2])]
    dig_plan.append((direction, distance))

print(dig_plan)


def calculate_vertices(dig_plan):
    directions = {"U": (0, 1), "D": (0, -1), "L": (-1, 0), "R": (1, 0)}
    vertices = [(0, 0)]  # Starting at (0, 0)
    current_position = [0, 0]

    for direction, distance in dig_plan:
        dx, dy = directions[direction]
        new_position = (
            current_position[0] + dx * distance,
            current_position[1] + dy * distance,
        )
        vertices.append(new_position)
        current_position = list(new_position)

    return vertices


def shoelace_formula(vertices):
    n = len(vertices)
    area = 0
    for i in range(n):
        j = (i + 1) % n
        area += vertices[i][0] * vertices[j][1]
        area -= vertices[j][0] * vertices[i][1]
    return abs(area) / 2


vertices = calculate_vertices(dig_plan)
lagoon_volume = shoelace_formula(vertices)
trench_length = sum(distance for _, distance in dig_plan)

# Half of the trench volume is considered in the total volume
half_trench_volume = trench_length / 2
adjusted_lagoon_volume = lagoon_volume + half_trench_volume + 1
print(int(adjusted_lagoon_volume))
