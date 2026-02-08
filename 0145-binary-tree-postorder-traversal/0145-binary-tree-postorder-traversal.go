/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func postorderTraversal(root *TreeNode) []int {
    var res []int
    
    var closure func(*TreeNode) 
    closure = func(node *TreeNode) {
        if node == nil {
            return
        }

        closure(node.Left)
        closure(node.Right)
        res = append(res, node.Val)
    }

    closure(root)

    return res
}