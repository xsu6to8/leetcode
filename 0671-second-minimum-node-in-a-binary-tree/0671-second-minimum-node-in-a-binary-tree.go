/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    var q []*TreeNode
    q = append(q, root)

    var values []int
    baseVal := root.Val
    for len(q) != 0 {
        currNode := q[0]
        q = q[1:]

        if currNode.Left != nil {
            if currNode.Left.Val <= baseVal {
                q = append(q, currNode.Left)
            }
            if currNode.Right.Val <= baseVal {
                q = append(q, currNode.Right)
            }
        }

        if currNode.Right != nil {
            if currNode.Left.Val > baseVal {
                values = append(values, currNode.Left.Val)
            }
            if currNode.Right.Val > baseVal {
                values = append(values, currNode.Right.Val)
            }
        }

    }

    if len(values) == 0 {
        return -1
    }

    if len(values) == 1 {
        return values[0]
    }

    slices.Sort(values)

    return values[0]
}