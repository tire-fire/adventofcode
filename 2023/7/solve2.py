#!/usr/bin/env python
import re
import numpy as np

#with open('example', 'r') as f:
with open("input", "r") as f:
    inp = [x.strip().split(' ') for x in f.readlines() if x.strip()]

# Define the hands and their bids
hands = inp

# Function to convert card to numerical value
def card_value(card):
    values = {'A': 14, 'K': 13, 'Q': 12, 'J': 1, 'T': 10}
    return values[card] if card in values else int(card)

def sort_hands(hand_info):
    hand, _ = hand_info
    # Convert each card in the hand to its numerical value, keeping the original order
    hand_values = [-1 * card_value(card) for card in hand]
    counts = {card: hand.count(card) for card in hand if card != 'J'}
    joker_count = hand.count('J')
    if counts:
        counts[max(counts, key=counts.get)] += joker_count
    else:
        counts = {card: hand.count(card) for card in hand}

    # Determine the type of hand for primary sorting
    if len(set(counts.values())) == 1 and list(counts.values())[0] == 5:
        hand_type = 1  # Five of a kind
    elif 4 in counts.values():
        hand_type = 2  # Four of a kind
    elif sorted(counts.values(), reverse=True) == [3, 2]:
        hand_type = 3  # Full house
    elif 3 in counts.values():
        hand_type = 4  # Three of a kind
    elif sorted(counts.values(), reverse=True) == [2, 2, 1]:
        hand_type = 5  # Two pair
    elif 2 in counts.values():
        hand_type = 6  # One pair
    else:
        hand_type = 7  # High card

    # Return a tuple for sorting: first by hand type, then by the card values in order
    return (hand_type, hand_values)


# Rank and sort the hands
ranked_hands = sorted([(hand, int(bid)) for hand, bid in hands], key=sort_hands, reverse=True)

# Calculate total winnings
total_winnings = sum(bid * (index + 1) for index, (_, bid) in enumerate(ranked_hands))


print(total_winnings)
