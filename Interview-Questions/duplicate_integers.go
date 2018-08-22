/*

1. Given an array of sequential (non-ordered) integers and a max number, return an array of the 2 integers which are duplicated
	func(int[] arr, int num) // the numbers 1-num will be in arr and each number in that sequence will either be in arr once or twice
	For example, func([2, 3, 4, 2, 1, 3], 4) -> [2, 3] // the numbers 1-4 are in the array, but only 2 and 3 are dupliated

	Challenge: Find O(n) solution without using a map/dictionary.

*/

package main

import (
	"math"
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
			fmt.Println("Error has occurred: \n", err.Error())
			return
		}
		if n <= 0 {
			fmt.Println("Only positive numbers allowed")
			return
		}
		if n > max {
			max = n
		}
		arr = append(arr, n)
	}

	duplicates := duplicateIntegers(arr, max)
	duplicatesUsingMap := duplicateIntegersUsingMap(arr)

	fmt.Println("\nThe duplicate numbers are: ")
	fmt.Println(duplicates)
	fmt.Println(duplicatesUsingMap)
}

func duplicateIntegersUsingMap(arr []int) []int {
	m := make(map[int]bool, 0)
	returnArr := make([]int, 0)
	for _, num := range arr {
		if _, exists := m[num]; exists {
			returnArr = append(returnArr, num)
		} else {
			m[num] = true
		}
	}
	return returnArr
}

func duplicateIntegers(arr []int, max int) []int {
	sum := 0
	product := 1

	for _, num := range arr {
		sum += num
		product *= num
	} // O(n)

	expectedSum := getExpectedSum(max) // O(n)
	expectedProduct := factorial(max) // O(n)

	sumDifference := sum - expectedSum
	productDifference := product / expectedProduct

	// a + b = sumDifference
	// a * b = productDifference
	// a = sumDifference - b
	// (sumDifference - b) * b = productDifference
	// b^2 - b*sumDiffernce + productDifference = 0

	b := (sumDifference + int(math.Sqrt(float64(sumDifference*sumDifference - (4 * productDifference)))))/2
	a := sumDifference - b

	returnArr := make([]int, 0)
	returnArr = append(returnArr, a)
	returnArr = append(returnArr, b)
	return returnArr
}

// returns the sum of numbers from 1 to n
func getExpectedSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// returns factorial of n
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}