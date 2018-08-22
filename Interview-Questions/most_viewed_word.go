/*

4. Given the following struct/object:

	Post {
		String text // is an arbitrary (but reasonable) number of lower-case, space-delimited strings
		int views // number of views that the entire text received
	}

	write a function that takes a list of Post objects and determines the most viewed word, defined as follows:
		For each 'text', each word attains the number of views for the entire Post
		For example, 
			Post {
				"hello world"
				10
			}
			Post {
				"hello universe"
				25
			}
			"hello" is the most viewed word with 35 views
			"world has 10 views
			"universe" has 25 views

	Challenge: Find O(n) solution.

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type Post struct {
	views int
	text string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	posts := make([]Post, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "DONE" {
			break
		}

		fields := strings.Fields(line)

		text := ""
		for i := 0; i < len(fields)-1; i++ {
			text += fields[i]
			text += " "
		}

		post := Post {
			text: text,
		}

		views, err := strconv.Atoi(fields[len(fields)-1])
		if err != nil {
			fmt.Println("Error has occurred: \n", err.Error())
			return
		}

		post.views = views
		posts = append(posts, post)
	}

	mostViewedWord, views := mostViewedWord(posts)

	fmt.Println(fmt.Sprintf("\nThe most viewed word is %s with %d views", mostViewedWord, views))
}

func mostViewedWord(posts []Post) (string, int) {
	mostViews := 0
	mostViewedWord := ""
	wordViews := make(map[string]int, 0)
	for _, post := range posts {
		words := strings.Fields(post.text)
		for _, word := range words {
			if numViews, exists := wordViews[word]; exists {
				wordViews[word] = numViews + post.views
			} else {
				wordViews[word] = post.views
			}
			if wordViews[word] > mostViews {
				mostViews = wordViews[word]
				mostViewedWord = word
			}
		}
	}
	return mostViewedWord, mostViews
}