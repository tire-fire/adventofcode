#!/usr/bin/env python
import re
import numpy as np

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [list(map(int, x.strip().split(" "))) for x in f.readlines() if x.strip()]


def extrapolate_next_value(history):
    # Calculate differences until all differences are zero
    differences = []
    while not all(v == 0 for v in history):
        differences.append(history)
        history = [history[i + 1] - history[i] for i in range(len(history) - 1)]

    # Add a zero to start the back extrapolation
    history.insert(0, 0)

    # Back extrapolate to find the next value
    for i in range(len(differences)):
        history.insert(0, differences[-i - 1][0] - history[0])

    # The first value in history is the extrapolated next value
    return history[0]


sum_extrapolated_values = sum(extrapolate_next_value(history) for history in inp)
print(sum_extrapolated_values)
