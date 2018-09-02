package main

import (
    "fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func InsertAll(arr []int) *TreeNode {
    var root *TreeNode
    for _, num := range arr {
        root = insert(root, num)
    }
    return root
}

func insert(root *TreeNode, num int) *TreeNode {
    if root == nil {
        return &TreeNode {
            Val: num,
        }
    } else {
        if num < root.Val {
            root.Left = insert(root.Left, num)
        } else {
            root.Right = insert(root.Right, num)
        }
    }
    return root
}

func PrintTree(root *TreeNode) {
    fmt.Println("")
    printTreeRecursive(root, 0)
}

func printTreeRecursive(root *TreeNode, level int) {
    fmt.Print(fmt.Sprintf("Row %d ", level))
    fmt.Println("Val: ", root.Val)
    if root.Left != nil {
        fmt.Print("Left ")
        printTreeRecursive(root.Left, level+1)
    }
    if root.Right != nil {
        fmt.Print("Right ")
        printTreeRecursive(root.Right, level+1)
    }
}