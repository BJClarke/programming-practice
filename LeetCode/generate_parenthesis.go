/*

Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

*/

package main

import (
    "bufio"
    "os"
    "fmt"
    "strconv"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    num := scanner.Text()
    n, err := strconv.Atoi(num)
    if err != nil {
        fmt.Println("Error has occurred: \n", err.Error())
        return
    }

    combinations := generateParenthesis(n)

    fmt.Println("\nThe possible combinations are: ")
    fmt.Println(combinations)
}

func generateParenthesis(n int) []string {
    solutions := new([]string)
    if n < 1 {
        return *solutions
    }
    str := ""
    for i := 0; i < n; i++ {
        str += "("
    }
    index := n
    for index > 0 {
        generateStrings(solutions, str, n)
        str = str[1:index] + ")"
        index--
    }
    return *solutions
}

func generateStrings(solutions *[]string, str string, n int) {
    if len(str) == 2*n {
        if isValid(str) {
            *solutions = append(*solutions, str)
        }
        return
    }
    
    generateStrings(solutions, str+"(", n)
    generateStrings(solutions, str+")", n)
}

func isValid(str string) bool {
    return valid(str, 0)
}

func valid(str string, numOpened int) bool {
    if len(str) == 0 {
        return numOpened == 0
    }
    
    if numOpened < 0 {
        return false
    }
    
    if str[0:1] == "(" {
        return valid(str[1:], numOpened+1)
    }
    return valid(str[1:], numOpened-1)
}

