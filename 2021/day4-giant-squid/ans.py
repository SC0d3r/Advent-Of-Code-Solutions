#!/usr/bin/env python3

import numpy as np


lines = []

with open("real-data.txt") as f:
    lines = f.readlines()

# getting the numbers
ns = list(map(int, lines[0].split(",")))

cursor = 1
def next_card(cursor,lines):
    rows = []
    cursor += 1 #skip one empyt line
    for i in range(5):
        rows.append(list(map(int, lines[cursor].split())))
        cursor += 1

    card = [row for row in rows]
    card = np.array(card)
    return card, cursor


cards = []
while cursor < len(lines):
    card, cursor = next_card(cursor, lines)
    cards.append(card)



def eq(ar1,ar2):
    for x in ar1:
        if x not in ar2:
            return False
    return True

def is_card_winner(card,seen_ns):
    #print("card",card)
    for row in card:
        if eq(row,seen_ns):
            return True
    return False

def calc_val(card,seen_ns,last_val):
    unseen_ns = []
    for row in card:
        for x in row:
            if x not in seen_ns:
                unseen_ns.append(x)
    print(sum(unseen_ns))
    print(last_val)
    return sum(unseen_ns) * last_val

seen = []
for x in ns:
    seen.append(x)
    for i, c in enumerate(cards):

        if is_card_winner(c, seen):
         print("card%s is winner" % (i + 1),calc_val(c,seen, x))
         exit(0)



print(seen)
