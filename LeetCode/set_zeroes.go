/*

Given a m x n matrix, if an element is 0, set its entire row and column to 0. Do it in-place.

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

    matrix := make([][]int, 0)

    for scanner.Scan() {
        line := scanner.Text()

        if line == "DONE" {
            break
        }

        arr := make([]int, 0)

        nums := strings.Fields(line)

        for _, num := range nums {
            n, err := strconv.Atoi(num)
            if err != nil {
                fmt.Println("Error has occured: \n", err.Error())
                return
            }  
            arr = append(arr, n)
        }

        matrix = append(matrix, arr)

        if len(arr) != len(matrix[0]) {
            fmt.Println("Error has occurred: The length of all rows should be equal")
            return
        }
    }

    setZeroes(matrix)

    fmt.Println("\nThe new matrix is: ")
    for _, row := range matrix {
    	fmt.Println(row)
    }
}

func setZeroes(matrix [][]int)  {
    zeroRows := make(map[int]bool, 0)
    zeroCols := make(map[int]bool, 0)
    
    for i, row := range matrix {
        for j, val := range row {
            if val == 0 {
                zeroRows[i] = true
                zeroCols[j] = true
            }
        }
    }
    
    for i, row := range matrix {
        for j := range row {
            if zeroRows[i] || zeroCols[j] {
                matrix[i][j] = 0
            }
        }
    }
}