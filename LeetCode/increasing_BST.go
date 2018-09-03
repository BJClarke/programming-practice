/*

Given a tree, rearrange the tree in in-order so that the leftmost node in the tree is now the root of the tree,
and every node has no left child and only 1 right child.

Example 1:
Input: [5,3,6,2,4,null,8,1,null,null,null,7,9]

       5
      / \
    3    6
   / \    \
  2   4    8
 /        / \ 
1        7   9

Output: [1,null,2,null,3,null,4,null,5,null,6,null,7,null,8,null,9]

 1
  \
   2
    \
     3
      \
       4
        \
         5
          \
           6
            \
             7
              \
               8
                \
                 9  

Note:

The number of nodes in the given tree will be between 1 and 100.
Each node will have a unique integer value from 0 to 1000.

*/

func increasingBST(root *TreeNode) *TreeNode {
    nodes := inorder(root, make([]int, 0))
    return insertAll(nodes)
}

func insertAll(arr []int) *TreeNode {
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
        root.Right = insert(root.Right, num)
    }
    return root
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