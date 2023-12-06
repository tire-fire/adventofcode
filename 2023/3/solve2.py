#!/usr/bin/env python
import re
from math import prod

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

def find_gears_and_calculate_sum(schematic):
    def in_bounds(x, y):
        return 0 <= x < len(schematic) and 0 <= y < len(schematic[x])

    def get_number_at(x, y):
        # Check if we are on a digit and find the full number it's a part of
        if not in_bounds(x, y) or not schematic[x][y].isdigit():
            return None
        num = ''
        # Move left to the start of the number
        while in_bounds(x, y - 1) and schematic[x][y - 1].isdigit():
            y -= 1
        # Move right to get the full number
        while in_bounds(x, y) and schematic[x][y].isdigit():
            num += schematic[x][y]
            y += 1
        return int(num)

    def find_adjacent_numbers(x, y):
        numbers = set()
        for dx in [-1, 0, 1]:
            for dy in [-1, 0, 1]:
                if dx == 0 and dy == 0:
                    continue
                num = get_number_at(x + dx, y + dy)
                if num is not None:
                    numbers.add(num)
        return numbers

    total_gear_ratio = 0

    for x, row in enumerate(schematic):
        for y, char in enumerate(row):
            if char == '*':
                adjacent_numbers = find_adjacent_numbers(x, y)
                if len(adjacent_numbers) == 2:
                    total_gear_ratio += prod(adjacent_numbers)

    return total_gear_ratio

# Example usage
schematic = [
    "467..114..",
    "...*......",
    "..35..633.",
    "......#...",
    "617*......",
    ".....+.58.",
    "..592.....",
    "......755.",
    "...$.*....",
    ".664.598.."
]
schematic = inp

print(find_gears_and_calculate_sum(schematic))
