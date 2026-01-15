/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    maxD := 0

    var cnt func(node *TreeNode, depth int)
    cnt = func(node *TreeNode, depth int) {
        // base condition
        if node == nil {
            return
        }

        if maxD < depth + 1 {
            maxD = depth + 1
        }

        cnt(node.Left, depth + 1)
        cnt(node.Right, depth + 1)
    }

    cnt(root, 0)

    return maxD
}