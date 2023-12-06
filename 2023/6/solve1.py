#!/usr/bin/env python
import re
import numpy as np

with open("input", "r") as f:
    inp = [list(map(int, x.strip().split(':').pop().split())) for x in f.readlines() if x.strip()]

races = list(zip(inp[0], inp[1]))
print(races)

def calculate_ways_to_win(time_limit, record_distance):
    ways_to_win = 0
    for hold_time in range(time_limit):
        speed = hold_time
        travel_time = time_limit - hold_time
        distance = speed * travel_time
        if distance > record_distance:
            ways_to_win += 1
    return ways_to_win


total_ways = np.prod([calculate_ways_to_win(*race) for race in races])

print(total_ways)
