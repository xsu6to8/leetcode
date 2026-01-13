/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    res := true
    
    var comp func(*TreeNode, *TreeNode)
    comp = func(curr1 *TreeNode, curr2 *TreeNode) {
        if !res {
            return
        }

        // Case 1: Both nil -> good
        if curr1 == nil && curr2 == nil {
            return
        }
        
        // Case 2: Only one nil -> wrong
        if curr1 == nil || curr2 == nil {
            res = false
            return
        }

        // Case 3: Compare Val
        if curr1.Val != curr2.Val {
            res = false
            return
        }

        comp(curr1.Left, curr2.Left)
        comp(curr1.Right, curr2.Right)
    }

    comp(p, q)
    return res
}