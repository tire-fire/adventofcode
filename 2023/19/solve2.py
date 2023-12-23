#!/usr/bin/env python
import re
import numpy as np
from itertools import combinations, product
from tqdm import tqdm
from heapq import heappush, heappop

# with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

workflow_pattern = r"(\w+)\{(.+?)\}"
rule_pattern = r"(\w+<\d+:\w+|\w+>\d+:\w+|.\w+:.\w+)"
part_pattern = r"\{x=(\d+),m=(\d+),a=(\d+),s=(\d+)\}"
workflows, parts = {}, []
for line in inp:
    if re.match(workflow_pattern, line):
        workflow_name, rules_str = re.match(workflow_pattern, line).groups()
        rules = [
            tuple(rule.split(":")) if ":" in rule else ("", rule)
            for rule in rules_str.split(",")
        ]
        workflows[workflow_name] = rules
    elif re.match(part_pattern, line):
        x, m, a, s = re.match(part_pattern, line).groups()
        parts.append({"x": int(x), "m": int(m), "a": int(a), "s": int(s)})

print(workflows)
print(parts)


def evaluation_loop(ranges, token):
    if token == "A":
        return len(ranges["x"]) * len(ranges["m"]) * len(ranges["a"]) * len(ranges["s"])
    if token == "R":
        return 0

    count = 0
    for condition, destination in workflows[token]:
        if condition:
            lhs, rhs = (
                condition.split("<") if "<" in condition else condition.split(">")
            )
            rhs = int(rhs)

            if "<" in condition:
                branch_true = set(filter(lambda x: x < rhs, ranges[lhs]))
                branch_false = set(filter(lambda x: x >= rhs, ranges[lhs]))
            else:
                branch_true = set(filter(lambda x: x > rhs, ranges[lhs]))
                branch_false = set(filter(lambda x: x <= rhs, ranges[lhs]))

            range_true = {
                key: (branch_true if key == lhs else val) for key, val in ranges.items()
            }
            range_false = {
                key: (branch_false if key == lhs else val)
                for key, val in ranges.items()
            }

            count += evaluation_loop(range_true, destination)
            ranges[lhs] = branch_false
        else:
            count += evaluation_loop(ranges, destination)

    return count


initial_ranges = {
    "x": set(range(1, 4001)),
    "m": set(range(1, 4001)),
    "a": set(range(1, 4001)),
    "s": set(range(1, 4001)),
}
total_accepted_combinations = evaluation_loop(initial_ranges, "in")
print(total_accepted_combinations)
