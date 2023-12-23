#!/usr/bin/env python
import re
import numpy as np
from collections import deque
from itertools import combinations
from tqdm import tqdm

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

steps = inp[0].split(",")


def holiday_ascii_string_helper(s):
    current_value = 0
    for char in s:
        ascii_code = ord(char)
        current_value += ascii_code
        current_value *= 17
        current_value %= 256
    return current_value


hash_results = [holiday_ascii_string_helper(step) for step in steps]


def hashmap_procedure(sequence):
    boxes = {i: [] for i in range(256)}

    for step in steps:
        label, operation = step.split("=" if "=" in step else "-")
        box_number = holiday_ascii_string_helper(label)

        # Removing a lens
        if operation == "":
            for lens in boxes[box_number]:
                if lens[0] == label:
                    boxes[box_number].remove(lens)
                    break

        # Adding or replacing a lens
        else:
            focal_length = int(operation)
            lens_found = False
            for i, lens in enumerate(boxes[box_number]):
                if lens[0] == label:
                    boxes[box_number][i] = (label, focal_length)
                    lens_found = True
                    break
            if not lens_found:
                boxes[box_number].append((label, focal_length))

    total_focusing_power = 0
    for box_number, lenses in boxes.items():
        for slot_number, (label, focal_length) in enumerate(lenses, start=1):
            focusing_power = (box_number + 1) * slot_number * focal_length
            total_focusing_power += focusing_power

    return total_focusing_power


total = hashmap_procedure(steps)
print(total)
