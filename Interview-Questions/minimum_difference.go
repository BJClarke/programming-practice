/*

2. Given an array of numbers, find the minimum difference between any two numbers in the array
	func(int[]) -> int
	For example, func([10, 7, 2, 5, 3]) -> 1 because the difference between 2 and 3 is 1

	Challenge: Finc O(n) solution.

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	arr := make([]int, 0)

	scanner.Scan()
	line := scanner.Text()

	nums := strings.Fields(line)

	max := 0

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error has occured: \n", err.Error())
			return
		}
		if n < 0 {
			fmt.Println("No negative numbers allowed")
			return
		}
		if n > max {
			max = n
		}
		arr = append(arr, n)
	}

	if len(arr) < 2 {
		fmt.Println("You must enter at least 2 numbers")
	}

	min := minimumDifference(arr)

	fmt.Println("\nThe minimum difference is: ")
	fmt.Println(min)
}

func minimumDifference(arr []int) int {
	max := getMax(arr)
	min := max

	buckets := make([]bool, max+1)
	for _, num := range arr {
		buckets[num] = true
	} // O(n)

	index := -1
	for i, isNum := range buckets {
		if isNum {
			if index != -1 {
				if i-index < min {
					min = i-index
				}
			}
			index = i
		}
	} // O(m)

	return min
}

func getMax(arr []int) int {
	max := arr[0]
	for _, num := range arr {
		if num > max {
			max = num
		}
	}
	return max
}