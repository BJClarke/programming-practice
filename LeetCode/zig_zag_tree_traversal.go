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

    result := zigzagLevelOrder(root)

    fmt.Println("\nThe zig zag traversal of the tree is: ")
    for _, level := range result {
        fmt.Println(level)
    }
}

func zigzagLevelOrder(root *TreeNode) [][]int {
    result := make([][]int, 0)
    if root == nil {
        return result
    }
    return zigzag([]*TreeNode{root}, result, true)
}

func zigzag(nodes []*TreeNode, v [][]int, leftToRight bool) [][]int {
    values := make([]int, 0)
    
    if len(nodes) == 0 {
        return v
    }
    
    children := make([]*TreeNode, 0)
    
    for _, node := range nodes {
        values = append(values, node.Val)
        if leftToRight {
            if node.Left != nil {
                children = append([]*TreeNode{node.Left}, children...)
            }
            if node.Right != nil {
                children = append([]*TreeNode{node.Right}, children...)
            }
        } else {
            if node.Right != nil {
                children = append([]*TreeNode{node.Right}, children...)
            }
            if node.Left != nil {
                children = append([]*TreeNode{node.Left}, children...)
            }
        }
    }
    v = append(v, values)
    return zigzag(children, v, !leftToRight)
}