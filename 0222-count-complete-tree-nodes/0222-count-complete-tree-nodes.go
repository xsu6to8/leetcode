/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    if root == nil {
        return 0
    }

    HLeft, HRight := 0, 0
    NodeL, NodeR := root, root

    for NodeL != nil {
        HLeft++
        NodeL = NodeL.Left
    }

    for NodeR != nil {
        HRight++
        NodeR = NodeR.Right
    }

    if HLeft == HRight {
        return (1 << HLeft) - 1
    }

    return 1 + countNodes(root.Left) + countNodes(root.Right)
}