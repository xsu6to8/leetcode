/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func hasPathSum(root *TreeNode, targetSum int) bool {
    // empty tree edge case
    if root == nil {
        return false
    }

    // arrive at 'leaf' 
    if root.Left == nil && root.Right == nil {
        // current val == remain of sum
        return root.Val == targetSum
    }

    // update remaining val
    remainingSum := targetSum - root.Val

    // if anything is [true]
    return hasPathSum(root.Left, remainingSum) || hasPathSum(root.Right, remainingSum)
}