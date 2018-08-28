/*

For the purposes of this challenge, we define a binary search tree to be a binary tree with the following properties:

The data value of every node in a node's left subtree is less than the data value of that node.
The data value of every node in a node's right subtree is greater than the data value of that node.
The data value of every node is distinct.

*/


public boolean checkBST(Node root) {
  
  if (root == null) {
    return true;
  }
  
  return checkTree(root.left, 0, root.data) && checkTree(root.right, root.data, Integer.MAX_VALUE);
}

public boolean checkTree(Node root, int left, int right) {
  
  if (root == null) {
    return true;
  }
  
  if (root.data < right && root.data > left) {
    // left child
    int leftNewLeft = left;
    int leftNewRight = root.data;
    
    // right child
    int rightnewLeft = root.data;
    int rightNewRight = right;
    
    return checkTree(root.left, leftNewLeft, leftNewRight) && checkTree(root.right, rightNewLeft, rightNewRight);
  }
  
  return false;
  
}