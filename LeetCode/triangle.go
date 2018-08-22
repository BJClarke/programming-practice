/*

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triangle

[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]

]

The minimum path sum from top to bottom is 11 (i.e., 2 + 3 + 5 + 1 = 11).

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

    triangle := make([][]int, 0)

    expectedLength := 1

    for scanner.Scan() {
        line := scanner.Text()

        if line == "DONE" {
            break
        }

        arr := make([]int, 0)

        nums := strings.Fields(line)

        if len(nums) != expectedLength {
            fmt.Println("Error has occurred: The length of the rows should begin at 1 and increase by 1")
            return
        }

        expectedLength++

        for _, num := range nums {
            n, err := strconv.Atoi(num)
            if err != nil {
                fmt.Println("Error has occured: \n", err.Error())
                return
            }  
            arr = append(arr, n)
        }
        triangle = append(triangle, arr)
    }

    minSum := minimumTotal(triangle)

    fmt.Println("\nThe minimum sum is: ")
    fmt.Println(minSum)
}


func minimumTotal(triangle [][]int) int {
    arr := make([]int, 0)
    arr = append(arr, triangle[0][0])
    return minTotal(triangle, arr, 1)
}

func minTotal(t [][]int, sums []int, row int) int {
    if row == len(t) {
        return getMin(sums)
    }
    newSums := make([]int, row+1)
    for i, num := range t[row] {
        if i > 0 && i < len(t[row])-1 {
            newSums[i] = min(sums[i-1], sums[i]) + num
        } else if i == 0 {
            newSums[i] = sums[0] + num
        } else {
            newSums[i] = sums[len(sums)-1] + num
        }
    }
    return minTotal(t, newSums, row+1)
}

func getMin(arr []int) int {
    min := arr[0]
    for _, num := range arr {
        if num < min {
            min = num
        }
    }
    return min
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}