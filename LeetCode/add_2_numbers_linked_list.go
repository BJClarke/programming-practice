/*

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	line1 := scanner.Text()
	num1, _ := strconv.Atoi(line1)

	scanner.Scan()
	line2 := scanner.Text()
	num2, _ := strconv.Atoi(line2)

	l1 := new(ListNode)
	list1 := l1

	l2 := new(ListNode)
	list2 := l2

	for num1 > 0 {
		n := num1 % 10
		list1.Val = n
		num1 /= 10
		if num1 > 0 {
			list1.Next = new(ListNode)
			list1 = list1.Next
		}
	}

	for num2 > 0 {
		n := num2 % 10
		list2.Val = n
		num2 /= 10
		if num2 > 0 {
			list2.Next = new(ListNode)
			list2 = list2.Next
		}
	}


	l3 := addTwoNumbers(l1, l2)

	factor := 1
	sum := 0
	for l3 != nil {
		sum += l3.Val*factor
		factor *= 10
		l3 = l3.Next
	}

	fmt.Println("\nThe sum of " + line1 + " and " + line2 + " is: ")
	fmt.Println(sum)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    head := new(ListNode)
	result := head
	oldCarry := 0
	newCarry := 0
	for l1 != nil && l2 != nil {
		sum := l1.Val + l2.Val + oldCarry
		if sum > 9 {
			sum = sum - 10
			newCarry = 1
		}
		result.Next = new(ListNode)
		result.Val = sum
		oldCarry = newCarry
		newCarry = 0

		l1 = l1.Next
		l2 = l2.Next

		if l1 != nil || l2 != nil {
			result = result.Next
		}
	}
	var l3 *ListNode
	if l1 != nil {
		l3 = l1
	} else if l2 != nil {
		l3 = l2
	}
	for l3 != nil {
		sum := l3.Val + oldCarry
		if sum > 9 {
			sum = sum - 10
			newCarry = 1
		}
		result.Next = new(ListNode)
		result.Val = sum
		oldCarry = newCarry
		newCarry = 0
		l3 = l3.Next
		if l3 != nil {
			result = result.Next
		}
	}

	if oldCarry != 0 {
		result.Next = new(ListNode)
		result = result.Next
		result.Val = oldCarry
	}

	result.Next = nil
	return head
}