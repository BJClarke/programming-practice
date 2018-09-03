/*

Given a binary search tree, write a function kthSmallest to find the kth smallest element in it.

Note: 
You may assume k is always valid, 1 ≤ k ≤ BST's total elements.


Example 1:

Input: root = [3,1,4,null,2], k = 1
   3
  / \
 1   4
  \
   2

Output: 1



Example 2:

Input: root = [5,3,6,2,4,null,null,1], k = 3
       5
      / \
     3   6
    / \
   2   4
  /
 1

Output: 3

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

    scanner.Scan()
    line = scanner.Text()
    k, err := strconv.Atoi(line)
    if err != nil {
        fmt.Println("Error has occurred: \n", err.Error())
        return
    }

    root := InsertAll(arr)

    PrintTree(root)

    kth := kthSmallestUsingInorder(root, k)

    str := getSuperlative(k, line)

    fmt.Println("\nThe " + str + " smallest element in the tree is: ")
    fmt.Println(kth)

    kth = kthSmallest(root, k)

    fmt.Println("\nThe " + str + " smallest element in the tree is: ")
    fmt.Println(kth)

}

func kthSmallest(root *TreeNode, k int) int {
    var solution int
    kth(root, &k, &solution)
    return solution
}

func kth(root *TreeNode, k *int, v *int) {
    if root == nil {
        return
    }
    
    kth(root.Left, k, v)
    
    *k = *k-1
    if *k == 0 {
        *v = root.Val
    }
    
    kth(root.Right, k, v)
}

func kthSmallestUsingInorder(root *TreeNode, k int) int {
    return inorder(root, make([]int, 0))[k-1]
}

func inorder(root *TreeNode, v []int) []int {
    if root == nil {
        return v
    }
    
    if root.Left != nil {
        v = inorder(root.Left, v)
    }
    
    v = append(v, root.Val)
    
    if root.Right != nil {
        v = inorder(root.Right, v)
    }
    
    return v
}

func getSuperlative(k int, line string) string {
    if k == 1 {
        return "\b"
    } else if k%10 == 1 {
        if k%100 == 11 {
            return line+"th"
        }
        return line+"st"
    } else if k%10 == 2 {
        if k%100 == 12 {
            return line+"th"
        }
        return line+"nd"
    } else if k%10 == 3 {
        if k%100 == 13 {
            return line+"th"
        }
        return line+"rd"
    }
    return line+"th"
}