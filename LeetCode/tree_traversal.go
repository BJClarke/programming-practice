/*

Given a binary tree, return the preorder, inorder, and postorder traversal of its nodes' values.

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

    root := InsertAll(arr)

    PrintTree(root)

    pre := preorderTraversal(root)
    in := inorderTraversal(root)
    post := postorderTraversal(root)

    fmt.Println("\nThe preorder traversal of the tree is: ")
    fmt.Println(pre)

    fmt.Println("\nThe inorder traversal of the tree is: ")
    fmt.Println(in)

    fmt.Println("\nThe postorder traversal of the tree is: ")
    fmt.Println(post)
}

func inorderTraversal(root *TreeNode) []int {
    var solution []int
    return inorder(root, solution)
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

func preorderTraversal(root *TreeNode) []int {
    var solution []int
    return preorder(root, solution)
}

func preorder(root *TreeNode, v []int) []int {
    if root == nil {
        return v
    }

    v = append(v, root.Val)
    
    if root.Left != nil {
        v = inorder(root.Left, v)
    }
    
    if root.Right != nil {
        v = inorder(root.Right, v)
    }
    
    return v
}

func postorderTraversal(root *TreeNode) []int {
    var solution []int
    return postorder(root, solution)
}

func postorder(root *TreeNode, v []int) []int {
    if root == nil {
        return v
    }
    
    if root.Left != nil {
        v = inorder(root.Left, v)
    }
    
    if root.Right != nil {
        v = inorder(root.Right, v)
    }

    v = append(v, root.Val)
    
    return v
}