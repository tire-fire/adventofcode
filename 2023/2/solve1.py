#!/usr/bin/env python
import re

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

def check_line(cubes):
    for number, color in cubes:
        number = int(number)
        if color == 'green' and number > 13:
            return False
        elif color == 'red' and number > 12:
            return False
        elif color == 'blue' and number > 14:
            return False

    return True

ans = 0
for line in inp:
    gamenum, line = line.split(':', 1)
    gamenum = int(gamenum.split(' ').pop())
    games = [x.strip() for x in line.split(';')]
    cubes = re.findall(r"(?:(\d+) (blue|green|red))[, ]*", line)
    print(cubes)
    if check_line(cubes):
        ans += gamenum

print(ans)
