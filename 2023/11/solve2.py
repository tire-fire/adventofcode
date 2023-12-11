#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from tqdm import tqdm

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]


def calculate_distance(galaxy_map, expansion_factor):
    # Convert galaxy map to a numpy array
    galaxy_array = np.array([list(row) for row in galaxy_map])

    # Identify rows and columns with galaxies
    rows_with_galaxies = np.any(galaxy_array == "#", axis=1)
    cols_with_galaxies = np.any(galaxy_array == "#", axis=0)

    # Find galaxy positions
    galaxies = np.argwhere(galaxy_array == "#")

    # Function to calculate the number of expanded rows and columns between two points
    def count_expanded_rows_and_cols(point1, point2):
        # Count expanded rows
        expanded_rows = np.sum(
            ~rows_with_galaxies[
                min(point1[0], point2[0]) : max(point1[0], point2[0] + 1)
            ]
        )
        # Count expanded columns
        expanded_cols = np.sum(
            ~cols_with_galaxies[
                min(point1[1], point2[1]) : max(point1[1], point2[1] + 1)
            ]
        )

        return expanded_rows, expanded_cols

    # Calculate the sum of expanded distances for all galaxy pairs
    total_expanded_distance = 0
    for i in range(len(galaxies)):
        for j in range(i + 1, len(galaxies)):
            expanded_rows, expanded_cols = count_expanded_rows_and_cols(
                galaxies[i], galaxies[j]
            )
            distance = (
                expanded_rows * (expansion_factor - 1)
                + abs(galaxies[i][0] - galaxies[j][0])
            ) + (
                expanded_cols * (expansion_factor - 1)
                + abs(galaxies[i][1] - galaxies[j][1])
            )
            total_expanded_distance += distance

    return total_expanded_distance


total = calculate_distance(inp, 1000000)

print(total)
