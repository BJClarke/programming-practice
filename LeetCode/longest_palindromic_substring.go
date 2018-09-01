/*

Given a string s, find the longest palindromic substring in s.
You may assume that the maximum length of s is 1000.

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
    str := scanner.Text()

    subStr := longestPalindrome(str)

    fmt.Println("\nThe longest palindromic substring is: ")
    fmt.Println(subStr)
}

func longestPalindrome(s string) string {
    if s == "" {
        return ""
    }
    
    begin := 0
    end := 0
    
    for i := 0; i < len(s)-1; i++ {
        // expand around a single char
        l1 := expandAroundCenter(s, i, i)
        if l1 > end-begin+1 {
            begin = i-(l1/2)
            end = i+(l1/2)
        }
        
        // expand around a pair of chars
        l2 := expandAroundCenter(s, i, i+1)
        if l2 > end-begin+1 {
            begin = i+1-(l2/2)
            end = i+(l2/2)
        }
    }
    return s[begin : end+1]
}

func expandAroundCenter(s string, left, right int) int {
    l := left
    r := right
    
    for l > -1 && r < len(s) && s[l] == s[r] {
        l--
        r++
    }
    
    return r-l-1
}