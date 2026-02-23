/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
    if root == nil {
        return nil
    }

    res := []string{}

    var toLeaf func(*TreeNode, string)
    toLeaf = func(node *TreeNode, path string) { 
        // base condition
        if node == nil {
            return
        }

        if node.Left == nil && node.Right == nil {
            newVal := strconv.Itoa(node.Val)
            if path == "" {
                path = newVal
            } else {
                path += "->" + newVal
            }

            res = append(res, path)
            return
        }

        newVal := strconv.Itoa(node.Val)
        if path == "" {
            path = newVal
        } else {
            path += "->" + newVal
        }

        toLeaf(node.Left, path)
        toLeaf(node.Right, path)
    }   

    toLeaf(root, "")

    return res
}
