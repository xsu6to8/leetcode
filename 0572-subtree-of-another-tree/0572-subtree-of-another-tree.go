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

    var preOrder func(node *TreeNode) bool
    preOrder = func(node *TreeNode) bool {
        result := false
        if node == nil {
            return false
        }

        if node.Val == subRoot.Val {
            result = isSame(node, subRoot)
            if result == true {
                return result
            }
        }

        left := preOrder(node.Left)
        if left == true {
            return left
        }
        right := preOrder(node.Right)
        if right == true {
            return right
        }

        return result
    }

    res := preOrder(root)

    return res
}