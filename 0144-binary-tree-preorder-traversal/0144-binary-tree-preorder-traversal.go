/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
    var res []int
    
    var closure func(*TreeNode) 
    closure = func(node *TreeNode) {
        if node == nil {
            return
        }

        res = append(res, node.Val)
        closure(node.Left)
        closure(node.Right)
    }

    closure(root)

    return res
}