#!/usr/bin/env python
import re

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

def sum_part_numbers(schematic):
    symbols = {'*', '#', '+', '$', '&', '%', '-', '/', '=', '@'}

    def in_bounds(x, y):
        return 0 <= x < len(schematic) and 0 <= y < len(schematic[x])

    def has_symbol_around(x, y):
        for dx in [-1, 0, 1]:
            for dy in [-1, 0, 1]:
                nx, ny = x + dx, y + dy
                if (dx != 0 or dy != 0) and in_bounds(nx, ny) and schematic[nx][ny] in symbols:
                    return True
        return False

    def find_numbers():
        for x, row in enumerate(schematic):
            num = ''
            for y, char in enumerate(row):
                if char.isdigit():
                    num += char
                if not char.isdigit() or y == len(row) - 1:
                    if num:
                        yield x, y - len(num), int(num)
                    num = ''

    def check_around(x, y, number):
        for i in range(len(str(number))):
            if has_symbol_around(x, y + i):
                return True

        return False

    total = 0
    for x, y, number in find_numbers():
        if check_around(x, y, number):
            total += number

    return total

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
print(schematic)

print(sum_part_numbers(schematic))
