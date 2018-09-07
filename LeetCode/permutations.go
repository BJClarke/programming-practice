/*

Given a collection of * distinct * integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    line := scanner.Text()

    nums := strings.Fields(line)
    arr := make([]int, 0)

    for _, num := range nums {
        n, err := strconv.Atoi(num)
        if err != nil {
            fmt.Println("Error has occurred: \n", err.Error())
            return
        }
        arr = append(arr, n)
    }

    permutations := permute(arr)

    fmt.Println("\nThe possible permutations are: ")
    fmt.Println(permutations)
}

func permute(nums []int) [][]int {
    solutions := new([][]int)
    
    if len(nums) < 1 {
        return *solutions
    }
    makePermutations(solutions, make([]int, 0), nums)
    return *solutions
}

func makePermutations(solutions *[][]int, arr []int, nums []int) {
    if len(nums) == 0 {
        *solutions = append(*solutions, arr)
        return
    }
    
    length := len(nums)
    for i := 0; i < length; i++ {
        newArr := append(arr, nums[i])
        subArr := subArray(nums, i)
        makePermutations(solutions, newArr, subArr)
    }
}

func subArray(a []int, index int) []int {
    arr := make([]int, 0)
    for i := 0; i < len(a); i++ {
        if i != index {
            arr = append(arr, a[i])
        }
    }
    return arr
}