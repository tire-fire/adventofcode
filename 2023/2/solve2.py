#!/usr/bin/env python
import re
import numpy as np

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

ans = 0
for line in inp:
    gamenum, line = line.split(':', 1)
    gamenum = int(gamenum.split(' ').pop())
    games = [x.strip() for x in line.split(';')]
    cubes = re.findall(r"(?:(\d+) (blue|green|red))[, ]*", line)
    cubes = sorted(cubes, key=lambda x: (x[1], int(x[0])))
    max_values = {}
    for number, color in cubes:
        if color not in max_values or int(number) > max_values[color]:
            max_values[color] = int(number)

    power = np.prod([int(x) for x in max_values.values()])
    ans += power


print(ans)
