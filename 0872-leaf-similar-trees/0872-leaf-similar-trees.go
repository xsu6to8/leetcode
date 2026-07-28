/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
    var sli1 []int
    var sli2 []int

    var dfs func(curr *TreeNode, values *[]int)
    dfs = func(curr *TreeNode, values *[]int) {
        if curr == nil {
            return
        }

        if curr.Left == nil && curr.Right == nil {
            *values = append(*values, curr.Val)
            return
        }

        dfs(curr.Left, values)
        dfs(curr.Right, values)
    }   

    dfs(root1, &sli1)
    dfs(root2, &sli2)

    if len(sli1) != len(sli2) {
        return false
    }

    for i := 0; i < len(sli1); i++ {
        if sli1[i] != sli2[i] {
            return false
        }
    }

    return true
}