/*

Given a string, replace occurrances of multiple consecutive characters with the repeated character followed by the number of times it is 	repeated in a row

	For example, func("aabbcc") -> "a2b2c2"
				 func("abbbbc") -> "ab4c"

*/



package main

import (
	"bufio"
	"os"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	str := scanner.Text()

	newStr := replaceConsecutiveOccurrences(str)

	fmt.Println("\nThe new string is: ")
	fmt.Println(newStr)
}

func replaceConsecutiveOccurrences(str string) string {
	count := 1
	newStr := ""
	r := ""

	for _, c := range str {
		if r != string(c) {
			newStr += r
			if count > 1 {
				newStr += strconv.Itoa(count)
			}
			count = 1
			r = string(c)
		} else {
			count++
		}
	}

	newStr += r

	if count > 1 {
		newStr += strconv.Itoa(count)
	}

	return newStr
}