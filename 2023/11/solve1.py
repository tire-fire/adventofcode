#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from itertools import combinations
from tqdm import tqdm
from collections import deque

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]


def expand_universe(galaxy_map):
    galaxy_array = np.array([list(row) for row in galaxy_map])
    empty_rows = np.all(galaxy_array == ".", axis=1)
    empty_cols = np.all(galaxy_array == ".", axis=0)

    expanded_rows = []
    for i, row in enumerate(galaxy_array):
        expanded_rows.append(row)
        if empty_rows[i]:
            expanded_rows.append(row)
    expanded_array = np.array(expanded_rows)

    expanded_array = np.column_stack(
        [
            expanded_array[:, i]
            if not empty_cols[i]
            else np.column_stack([expanded_array[:, i], expanded_array[:, i]])
            for i in range(expanded_array.shape[1])
        ]
    )
    return expanded_array


expanded_universe = expand_universe(inp)

# Find all galaxies
galaxies = np.argwhere(expanded_universe == "#")

# Calculate the Manhattan distance between all pairs of galaxies
total = sum(
    abs(g1[0] - g2[0]) + abs(g1[1] - g2[1]) for g1, g2 in combinations(galaxies, 2)
)
print(total)
