/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
    // sum은 하위에서부터 더해가며 갱신
    sum := 0
    var traverse func(node *TreeNode) int 
    traverse = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        
        // post-order traverse
        leftSum := traverse(node.Left)
        rightSum := traverse(node.Right)

        gap := int(math.Abs(float64(leftSum - rightSum)))
        sum += gap

        return node.Val + leftSum + rightSum
    }

    _ = traverse(root)

    return sum
}