/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// // method1 :  Recursive
// func preorder(root *Node) []int {
//     var res []int
//     var traverse func(node *Node)
//     traverse = func(node *Node) {
//         if node == nil {
//             return
//         }
//         res = append(res, node.Val)

//         for _, childNode := range node.Children {
//             traverse(childNode)
//         }
//     }

//     traverse(root)

//     return res
// }

// method2 : iteratively
func preorder(root *Node) []int {
    var res []int
    if root == nil {
        return res
    }

    stack := []*Node{root}
    for len(stack) > 0 {
        node := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        res = append(res, node.Val)

        for i := len(node.Children) - 1; i >= 0; i-- {
            stack = append(stack, node.Children[i])
        }
    }
    return res
}