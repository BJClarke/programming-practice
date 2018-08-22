/*

3. Construct a BSTNode object/struct and complete the following function:
	func search(BSTNode, int) -> boolean // this function takes the root of a BST and returns true if the given 'int' exists in the tree

	Challenge: Finc O(log(n)) solution.

*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"fmt"
)

type BSTNode struct {
	val int
	left *BSTNode
	right *BSTNode
}

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

	scanner.Scan()
	numStr := scanner.Text()
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Println("Error has occurred: \n", err.Error())
		return
	}

	root := insertAll(arr)

	printTree(root)

	numExists := search(root, num)

	if numExists {
		fmt.Println(fmt.Sprintf("\nThe number %d exists in the tree", num))
	} else {
		fmt.Println(fmt.Sprintf("\nThe number %d does NOT exist in the tree", num))
	}
	
}

func insertAll(arr []int) *BSTNode {
	var root *BSTNode
	for _, num := range arr {
		root = insert(root, num)
	}
	return root
}

func insert(root *BSTNode, num int) *BSTNode {
	if root == nil {
		return &BSTNode {
			val: num,
		}
	} else {
		if num < root.val {
			root.left = insert(root.left, num)
		} else {
			root.right = insert(root.right, num)
		}
	}
	return root
}

func printTree(root *BSTNode) {
	fmt.Println("")
	printTreeRecursive(root, 0)
}

func printTreeRecursive(root *BSTNode, level int) {
	fmt.Print(fmt.Sprintf("Row %d ", level))
	fmt.Println("Val: ", root.val)
	if root.left != nil {
		fmt.Print("Left ")
		printTreeRecursive(root.left, level+1)
	}
	if root.right != nil {
		fmt.Print("Right ")
		printTreeRecursive(root.right, level+1)
	}
}

func search(root *BSTNode, num int) bool {
	if root == nil {
		return false
	}

	if root.val == num {
		return true
	}

	if num < root.val {
		if root.left == nil {
			return false
		}
		return search(root.left, num)
	}
	if root.right == nil {
		return false
	}
	return search(root.right, num)
}