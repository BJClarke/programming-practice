/*

Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

2 -> a, b, c
3 -> d, e, f
4 -> g, h, i
5 -> j, k, l
6 -> m, n, o
7 -> p, q, r, s
8 -> t, u, v
9 -> w, x, y, z


Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].

*/

package main

import (
    "bufio"
    "os"
    "fmt"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    digits := scanner.Text()

    combinations := letterCombinations(digits)

    fmt.Println("\nThe possible combinations are: ")
    fmt.Println(combinations)
}


func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return make([]string, 0)
    }
    m := make(map[string][]string, 0)
    m["2"] = []string{"a", "b", "c"}
    m["3"] = []string{"d", "e", "f"}
    m["4"] = []string{"g", "h", "i"}
    m["5"] = []string{"j", "k", "l"}
    m["6"] = []string{"m", "n", "o"}
    m["7"] = []string{"p", "q", "r", "s"}
    m["8"] = []string{"t", "u", "v"}
    m["9"] = []string{"w", "x", "y", "z"}
    
    combinations := make([]string, 0)
    
    for i := 0; i < len(m[digits[0:1]]); i++ {
        combinations = append(combinations, m[digits[0:1]][i])
    }
    
    for i := 1; i < len(digits); i++ {
        newCombinations := make([]string, 0)
        for _, combo := range combinations {
            for k := 0; k < len(m[digits[i:i+1]]); k++ {
                newCombinations = append(newCombinations, combo+m[digits[i:i+1]][k])
            }
        }
        combinations = newCombinations
    }
    
    return combinations
}