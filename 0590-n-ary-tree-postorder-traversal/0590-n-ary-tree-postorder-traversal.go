/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    var res []int
    var traverse func(node *Node)
    traverse = func(node *Node) {
        if node == nil {
            return
        }
        for _, childNode := range node.Children {
            traverse(childNode)
        }

        res = append(res, node.Val)
    }

    traverse(root)

    return res
}