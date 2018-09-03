/*

We have an array A of non-negative integers.

For every (contiguous) subarray B = [A[i], A[i+1], ..., A[j]] (with i <= j), 
we take the bitwise OR of all the elements in B, obtaining a result A[i] | A[i+1] | ... | A[j].

Return the number of possible results.  (Results that occur more than once are only counted once in the final answer.)


Example 1:

Input: [0]
Output: 1
Explanation: 
There is only one possible result: 0.


Example 2:

Input: [1,1,2]
Output: 3
Explanation: 
The possible subarrays are [1], [1], [2], [1, 1], [1, 2], [1, 1, 2].
These yield the results 1, 1, 2, 1, 3, 3.
There are 3 unique values, so the answer is 3.


Example 3:

Input: [1,2,4]
Output: 6
Explanation: 
The possible results are 1, 2, 3, 4, 6, and 7.

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

    numUniqueValues := subarrayBitwiseORs(arr)

    fmt.Println("\nThe number of unique values is: ")
    fmt.Println(numUniqueValues)

}

func subarrayBitwiseORs(A []int) int {
    m := make(map[int]bool)
    temp1 := make(map[int]bool)
    temp2 := make(map[int]bool)
    for _, num := range A {
        temp2 = make(map[int]bool)
        temp2[num] = true
        for val := range temp1 {
            temp2[val|num] = true
        }
        temp1 = temp2
        for val := range temp1 {
            m[val] = true
        }
    }
    return len(m)
}