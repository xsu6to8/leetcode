/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    var res []int
    var traverse func(node *Node)
    traverse = func(node *Node) {
        if node == nil {
            return
        }
        res = append(res, node.Val)

        for _, childNode := range node.Children {
            traverse(childNode)
        }
    }

    traverse(root)

    return res
}