/*

Given a string, find the length of the longest substring without repeating characters.

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
    str := scanner.Text()

    length, subStr := lengthOfLongestSubstring(str)

    fmt.Println("\nThe longest substring without repeating characters is: ")
    fmt.Println("\"" + subStr + "\" which has length " + strconv.Itoa(length))
}

func lengthOfLongestSubstring(s string) (int, string) {
    m := make(map[rune]int, 0)
    begin := 0
    longest := 0
    end := 0
    subStr := ""
    var r rune
    
    for end, r = range s {
        if index := m[r]; index > begin {
            begin = index
        }
        m[r] = end + 1
        if current := end - begin + 1; current > longest {
            longest = current
            subStr = s[begin:end+1]
        }
    }
    
    return longest, subStr
}