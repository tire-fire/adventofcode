#!/usr/bin/env python
import re
import numpy as np
from itertools import combinations
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


def evaluate_condition(part, condition):
    # print(f"evaluating {part = } {condition = }")
    if condition == "":
        return True
    if "<" in condition:
        prop, value = condition.split("<")
        return part[prop] < int(value)
    elif ">" in condition:
        prop, value = condition.split(">")
        return part[prop] > int(value)


def process_part(part, current_workflow):
    for condition, destination in workflows[current_workflow]:
        if evaluate_condition(part, condition):
            if destination in ["A", "R"]:
                return destination
            return process_part(part, destination)
    return "R"


total_accepted_ratings = 0
for part in parts:
    result = process_part(part, "in")
    if result == "A":
        total_accepted_ratings += sum(part.values())

print(total_accepted_ratings)
