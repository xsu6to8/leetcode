/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    sum := 0

    var findLeaf func(curr *TreeNode, isLeft bool) 
    findLeaf = func(curr *TreeNode, isLeft bool) {
        if curr == nil {
            return
        }
        
        // base condition
        if curr.Left == nil && curr.Right == nil {
            if isLeft {
                sum += curr.Val
            }
            return
        }

        findLeaf(curr.Left, true)
        findLeaf(curr.Right, false)
    }

    findLeaf(root, false)

    return sum
}