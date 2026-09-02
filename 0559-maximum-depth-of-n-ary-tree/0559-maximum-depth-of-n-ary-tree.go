/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    maxD := 0

    var search func(node *Node, depth int) 
    search = func(node *Node, depth int) {
        if node == nil {
            return
        }

        if depth > maxD {
            maxD = depth
        }

        for _, child := range node.Children {
            if child != nil {
                search(child, depth+1)
            }
        }
    }

    search(root, 1)

    return maxD
}