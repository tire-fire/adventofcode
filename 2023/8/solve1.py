#!/usr/bin/env python
import re
import numpy as np

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

directions = [x for x in inp.pop(0)]
network = {}
for line in inp:
    line = re.findall(r"(\w+) = \((\w+), (\w+)\)", line)[0]
    network[line[0]] = (line[1], line[2])

def navigate_network(instructions, network, start_node, end_node):
    current_node = start_node
    steps = 0
    instruction_length = len(instructions)

    while current_node != end_node:
        # Get the next instruction, cycling through the list as needed
        direction = instructions[steps % instruction_length]

        # Move to the next node based on the instruction
        if direction == 'L':
            current_node = network[current_node][0]
        else:  # direction is 'R'
            current_node = network[current_node][1]

        steps += 1

    return steps

steps = navigate_network(directions, network, "AAA", "ZZZ")
print(steps)
