#!/usr/bin/env python2
import argparse

parser = argparse.ArgumentParser(description="Determines the number of hands that P1 (the first 3 numbers in each list of 6 numbers) wins")
parser.add_argument('input_file', type=str, nargs='?', help='the file containing test cases')
args = parser.parse_args()

hands = []

def find_num_hands_P1_wins():
	with open(args.input_file, 'r') as f:
		for line in f:
			hands.append(line)

	num_hands_P1_wins = 0
	
	for hand in hands:
		cards = hand.split(" ")
		p1 = []
		p1.append(cards[0])
		p1.append(cards[1])
		p1.append(cards[2])

		p2 = []
		p2.append(cards[3])
		p2.append(cards[4])
		p2.append(cards[5])


		if isP1Winner(p1, p2):
			num_hands_P1_wins = num_hands_P1_wins+1

	print num_hands_P1_wins

# hand value is 2 if 3 of a kind
#				1 if pair
#				0 if high card
def isP1Winner(p1, p2):
	p1_value = hand_value(p1)
	p2_value = hand_value(p2)
	if p1_value > p2_value:
		return True
	if p2_value > p1_value:
		return False

	if p1_value == 2:
		return p1[0] > p2[0]
	if p1_value == 1:
		return higher_pair(p1, p2)
	return higher_card(p1, p2)

def hand_value(hand):
	if hand[0] == hand[1] or hand[0] == hand[2] or hand[1] == hand[2]:
		if hand[0] == hand[1] and hand[0] == hand[2]:
			return 2
		return 1
	return 0

def higher_pair(p1, p2):
	p1_pair_value = 0
	p1_extra_card = 0

	p2_pair_value = 0
	p2_extra_card = 0

	if p1[0] == p1[1]:
		p1_pair_value = p1[0]
		p1_extra_card = p1[2]
	elif p1[1] == p1[2]:
		p1_pair_value = p1[1]
		p1_extra_card = p1[0]
	else:
		p1_pair_value = p1[2]
		p1_extra_card = p1[1]

	if p2[0] == p2[1]:
		p2_pair_value = p2[0]
		p2_extra_card = p2[2]
	elif p1[1] == p1[2]:
		p2_pair_value = p2[1]
		p2_extra_card = p2[0]
	else:
		p2_pair_value = p2[2]
		p2_extra_card = p2[1]

	if p1_pair_value > p2_pair_value:
		return True
	if p2_pair_value > p1_pair_value:
		return False

	return p1_extra_card > p2_extra_card

def higher_card(p1, p2):
	p1.sort()
	p2.sort()

	if p1[2] > p2[2]:
		return True
	if p2[2] > p1[2]:
		return False
	if p1[1] > p2[1]:
		return True
	if p2[1] > p1[1]:
		return False
	return p1[0] > p2[0]


if __name__ == '__main__': 
	find_num_hands_P1_wins()