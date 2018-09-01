/*

Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0?

Find all * unique * triplets in the array which gives the sum of zero.

Classic 3sum problem

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
	"sort"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	arr := make([]int, 0)

	scanner.Scan()
	line := scanner.Text()

	nums := strings.Fields(line)

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error has occurred: \n", err.Error())
			return
		}
		arr = append(arr, n)
	}

	arrs := threeSum(arr)

	fmt.Println("\nThe unique triplets are: ")
	for _, arr := range arrs {
		fmt.Println(arr)
	}
}

func threeSum(nums []int) [][]int {
    var solutions [][]int
    
	sort.Ints(nums)

    for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			return solutions
		}

        target := 0-nums[i]
        low := i+1
        high := len(nums)-1
        if i == 0 || nums[i] != nums[i-1] {
            for low < high {
                if nums[low] + nums[high] == target {
                    solutions = append(solutions, []int{nums[low], nums[high], nums[i]})
                    for low < high && nums[low] == nums[low+1] {
                        low++
                    }
                    for low < high && nums[high] == nums[high-1] {
                        high--
                    }
                    low++
                    high--
                } else if nums[low] + nums[high] < target {
                    low++
                } else {
                    high--
                }
            }
        }
    }
    
    return solutions
}