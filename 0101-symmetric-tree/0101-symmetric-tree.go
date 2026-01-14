/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return isMirror(root.Left, root.Right)
}

func isMirror(node1 *TreeNode, node2 *TreeNode) bool {
    // Base Condition
    if node1 == nil && node2 == nil {
        return true
    }
    if node1 == nil || node2 == nil {
        return false
    }

    // Different Vals -> false
    if node1.Val != node2.Val {
        return false
    }

    // Recursive 
    return isMirror(node1.Left, node2.Right) && isMirror(node1.Right, node2.Left)
}