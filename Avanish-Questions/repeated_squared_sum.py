#!/usr/bin/env python2
import argparse

parser = argparse.ArgumentParser(description="Determines the number of numbers less than or equal to n where the repeated squared sum of the digits equals 1")
parser.add_argument('n', type=int, nargs='?', help='the maximum number to check')
args = parser.parse_args()

def find_num_squared_sums():
	count = 0
	for x in range(args.n+1):
		if squared_sum(x):
			count = count+1
	print "There are %d such numbers" % count

def squared_sum(x):
	while x > 9:
		x = get_squared_sum(x)

	return x == 1

def get_squared_sum(x):
	sum = 0
	while x != 0:
		sum = sum + (x%10)**2
		x=x/10
	return sum

if __name__ == '__main__': 
	find_num_squared_sums()