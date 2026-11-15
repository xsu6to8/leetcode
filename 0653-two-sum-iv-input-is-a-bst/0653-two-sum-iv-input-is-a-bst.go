/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTarget(root *TreeNode, k int) bool {
    
    // BST -> 중위순회 == 오름차순
    var orderedNums []int
    var inorder func(node *TreeNode) 
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }

        inorder(node.Left)

        orderedNums = append(orderedNums, node.Val)

        inorder(node.Right)
    }

    inorder(root)

    l, r := 0, len(orderedNums)-1
    for l < r {
        lVal := orderedNums[l]
        rVal := orderedNums[r]

        currSum := lVal + rVal
        if currSum == k {
            return true
        } else if currSum < k {
            l++
        } else if currSum > k {
            r--
        }
    } 

    return false
}