#!/usr/bin/env python
import re

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

cards = [x.split(':')[1] for x in inp]
cards = [line.split('|') for line in cards]
cards = [(set(map(int, card[0].split())), set(map(int, card[1].split()))) for card in cards]

# Now, let's calculate the points for each card
total_points = 0
for winning_numbers, your_numbers in cards:
    matches = winning_numbers.intersection(your_numbers)
    if matches:
        # Calculate points: 1 point for the first match, doubled for each additional match
        points = 2 ** (len(matches) - 1)
        total_points += points

print(total_points)
