/*

Given an unsorted array return whether an increasing subsequence of length 3 exists or not in the array.

Return true if there exists i, j, k 
such that arr[i] < arr[j] < arr[k] given 0 ≤ i < j < k ≤ n-1 else return false.

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

	for _, num := range nums {
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Println("Error has occurred: \n", err.Error())
			return
		}
		arr = append(arr, n)
	}

	isIncreasingTriplet := increasingTriplet(arr)

	if isIncreasingTriplet {
		fmt.Println("\nThere IS an increasing triplet subsequence")
	} else {
		fmt.Println("\nThere is NOT an increasing triplet subsequence")
	} 
}

func increasingTriplet(nums []int) bool {
    
    if len(nums) < 3 {
        return false
    }
    
    isPairSubsequence := false
    
    var b int
    
    a := nums[0]
    newMin := a
    
    for i := 1; i < len(nums); i++ {
        num := nums[i]
        
        if isPairSubsequence {
            if num > b {
                return true
            }
            if num < a && i < len(nums)-2 {
                newMin = num
            } else if num < b && num > a {
                b = num
            } else if num < b && num > newMin {
                a = newMin
                b = num
            }
        } else {
            if num < a && i < len(nums)-2 {
                a = num
            } else if num > a {
                b = num
                isPairSubsequence = true
            }
        }
    }
    return false
}