/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


//  BST = Up-Down game 처럼 생각
func searchBST(root *TreeNode, val int) *TreeNode {
    curr := root

    for curr != nil && curr.Val != val {
        if curr.Val > val {
            curr = curr.Left
        } else if curr.Val < val {
            curr = curr.Right
        } else {
            break
        }
    }

    if curr == nil {
        return nil
    }

    return curr
}