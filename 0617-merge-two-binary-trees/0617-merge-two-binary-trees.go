/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
    
    var traverse func(node1, node2 *TreeNode) *TreeNode
    traverse = func(node1, node2 *TreeNode) *TreeNode {
        if node1 == nil && node2 != nil {
            return node2
        }
        if node1 != nil && node2 == nil {
            return node1
        }
        if node1 == nil && node2 == nil {
            return nil
        }

        //  pre-order
        node1.Val = node1.Val + node2.Val

        node1.Left = traverse(node1.Left, node2.Left)
        node1.Right = traverse(node1.Right, node2.Right)

        return node1
    }

    return traverse(root1, root2)
}