/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    // Base condition
    if root == nil {
        return nil
    }

    // Go의 다중 할당을 이용해 스왑과 재귀 호출을 한 번에 처리
    root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)

    return root
}