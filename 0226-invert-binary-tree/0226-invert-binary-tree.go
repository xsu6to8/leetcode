/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    invert(root)

    return root
}

func invert(node *TreeNode) {
    // base condition : 리프는 그냥 탈출
    if node == nil {
        return
    }

    tmpNode := node.Left
    node.Left = node.Right
    node.Right = tmpNode

    invert(node.Left)
    invert(node.Right)
}