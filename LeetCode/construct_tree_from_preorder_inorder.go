/*

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
You may assume that duplicates do not exist in the tree.

For example, given

preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7

*/

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }
    
    return buildTreeHelper(0, 0, len(inorder)-1, preorder, inorder)
}

func buildTreeHelper(preorderStart, inorderStart, inorderEnd int, preorder, inorder []int) *TreeNode {
    if preorderStart > len(preorder)-1 || inorderStart > inorderEnd {
        return nil
    }
    
    root := new(TreeNode)
    root.Val = preorder[preorderStart]
    rootInIndex := 0
    for inorder[rootInIndex] != root.Val {
        rootInIndex++
    }
    root.Left = buildTreeHelper(preorderStart+1, inorderStart, rootInIndex-1, preorder, inorder)
    root.Right = buildTreeHelper(preorderStart + rootInIndex - inorderStart + 1, rootInIndex+1, inorderEnd, preorder, inorder)
    return root
}