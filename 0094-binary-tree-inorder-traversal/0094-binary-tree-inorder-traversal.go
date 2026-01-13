/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var res []int

    var dfs func(*TreeNode)     //  Go -> func [declaration] FIRST!!!!
     
    dfs = func(node *TreeNode) {
        if node == nil {
            return
        }
        
        dfs(node.Left)             
        res = append(res, node.Val)
        dfs(node.Right)           
    }
    
    dfs(root)
    return res
}