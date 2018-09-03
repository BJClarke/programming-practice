/*

Given a binary tree

Populate each next pointer to point to its next right node. 
If there is no next right node, the next pointer should be set to NULL.

Initially, all next pointers are set to NULL.

Note:

You may only use constant extra space.
Recursive approach is fine, implicit stack space does not count as extra space for this problem.
You may assume that it is a perfect binary tree (ie, all leaves are at the same level, and every parent has two children).


Example:

Given the following perfect binary tree,

     1
   /  \
  2    3
 / \  / \
4  5  6  7

After calling your function, the tree should look like:

     1 -> NULL
   /  \
  2 -> 3 -> NULL
 / \  / \
4->5->6->7 -> NULL

*/

public class Solution {
    public void connect(TreeLinkNode root) {
        if (root == null) {
            return;
        }
        Queue<TreeLinkNode> nodes = new LinkedList<TreeLinkNode>();
        nodes.add(root);
        connectLevels(nodes);
    }

    public void connectLevels(Queue<TreeLinkNode> nodes) {
        if (nodes.size() == 0) {
            return;
        }
    
        Queue<TreeLinkNode> children = new LinkedList<TreeLinkNode>();
    
        while (!nodes.isEmpty()) {
            TreeLinkNode node = nodes.poll();
            if (!nodes.isEmpty()) {
                node.next = nodes.peek();
            }
            if (node.left != null) {
                children.add(node.left);
            }
            if (node.right != null) {
                children.add(node.right);
            }
        }
        connectLevels(children);
    }

    class TreeLinkNode {
        int val;
        TreeLinkNode left, right, next;
        TreeLinkNode(int x) { val = x; }
    }
}