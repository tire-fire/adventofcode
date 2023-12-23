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

total = sum(hash_results)
print(total)
