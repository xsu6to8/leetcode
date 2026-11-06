/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
    var isSame func(node1 *TreeNode, node2 *TreeNode) bool
    isSame = func(node1 *TreeNode, node2 *TreeNode) bool {        
        if node1 == nil && node2 == nil {
            return true
        }
        if node1 == nil && node2 != nil {
            return false
        }
        if node1 != nil && node2 == nil {
            return false
        }
        
        if node1.Val != node2.Val {
            return false
        }
        
        resLeft := isSame(node1.Left, node2.Left)
        resRight := isSame(node1.Right, node2.Right)

        return resLeft && resRight
    }

    var inOrder func(node *TreeNode) bool
    inOrder = func(node *TreeNode) bool {
        result := false
        if node == nil {
            return false
        }

        if node.Val == subRoot.Val {
            result = isSame(node, subRoot)
        }

        left := inOrder(node.Left)
        if left == true {
            return left
        }
        right := inOrder(node.Right)
        if right == true {
            return right
        }

        return result
    }

    res := inOrder(root)

    return res
}