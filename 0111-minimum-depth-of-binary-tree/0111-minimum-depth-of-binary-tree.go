/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    return helper(root)
}

func helper(node *TreeNode) int {
    if node == nil {
        return math.MaxInt64
    }

    if node.Left == nil && node.Right == nil {
        return 1
    }

    left := helper(node.Left)
    right := helper(node.Right)
    
    return min(left, right) + 1
}