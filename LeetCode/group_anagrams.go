/*

Given an array of lowercase strings, group anagrams together

Return an array of arrays of the original strings.

*/

package main

import (
    "bufio"
    "os"
    "sort"
    "fmt"
)

type ByRune []rune

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    strs := make([]string, 0)

    for scanner.Scan() {
        line := scanner.Text()

        if line == "DONE" {
            break
        }

        str := scanner.Text()
        strs = append(strs, str)
    }

    groups := groupAnagrams(strs)

    fmt.Println("\nThe groups are: ")
    for _, group := range groups {
    	fmt.Println(group)
    }
}

func groupAnagrams(strs []string) [][]string {
    arrIndex := make(map[string]int, 0)
    
    var solutions [][]string
    
    for _, str := range strs {
        sortedStr := SortStringByCharacter(str)
        if index, exists := arrIndex[sortedStr]; exists {
            solutions[index] = append(solutions[index], str)
        } else {
            arrIndex[sortedStr] = len(solutions)
            solutions = append(solutions, []string{str})
        }
    }
    
    return solutions
}

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func StringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}
	return r
}

func SortStringByCharacter(s string) string {
	var r ByRune = StringToRuneSlice(s)
	sort.Sort(r)
	return string(r)
}