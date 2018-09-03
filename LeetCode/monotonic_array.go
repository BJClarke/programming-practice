/*

An array is monotonic if it is either monotone increasing or monotone decreasing.

An array A is monotone increasing if for all i <= j, A[i] <= A[j].  
An array A is monotone decreasing if for all i <= j, A[i] >= A[j].

Return true if and only if the given array A is monotonic.

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

    isMonotone := isMonotonic(arr)

    if isMonotone {
        fmt.Println("\nThe array IS monotonic.")
    } else {
        fmt.Println("\nThe array is NOT monotonic.")
    }
}

func isMonotonic(A []int) bool {
    if len(A) < 2 {
        return true
    }
    
    for i := 1; i < len(A); i++ {
        if A[i] > A[i-1] {
            return isIncreasing(A)
        }
        if A[i] < A[i-1] {
            return isDecreasing(A)
        }
    }
    return true
}

func isIncreasing(A []int) bool {
    for i := 1; i < len(A); i++ {
        if A[i] < A[i-1] {
            return false
        }
    }
    return true
}

func isDecreasing(A []int) bool {
    for i := 1; i < len(A); i++ {
        if A[i] > A[i-1] {
            return false
        }
    }
    return true
}