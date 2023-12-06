#!/usr/bin/env python
import re

with open("input", "r") as f:
    inp = [x.strip() for x in f.readlines() if x.strip()]

cards = [x.split(':')[1] for x in inp]
cards = [line.split('|') for line in cards]
cards = [(set(map(int, card[0].split())), set(map(int, card[1].split()))) for card in cards]

def calculate_total_cards(cards):
    # Initialize a list to keep track of the number of copies of each card
    card_copies = [1] * len(cards)  # Start with 1 copy of each original card

    # Iterate over each card
    for i, (winning_numbers, your_numbers) in enumerate(cards):
        matches = winning_numbers.intersection(your_numbers)
        num_matches = len(matches)

        # For each match, add a copy of the next card
        for j in range(i + 1, min(i + 1 + num_matches, len(cards))):
            card_copies[j] += card_copies[i]

    # Sum all the copies
    return sum(card_copies)

# Calculate the total number of cards including originals and copies
total_cards = calculate_total_cards(cards)
print(total_cards)
