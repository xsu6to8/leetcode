/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func diameterOfBinaryTree(root *TreeNode) int {
    maxDiameter := 0

    var postOrder func(node *TreeNode) int
    postOrder = func(node *TreeNode) int {
        // 1. nil이면 깊이 0 반환
        if node == nil {
            return 0 
        }
        
        // 2. 자식들의 깊이를 먼저 받아옴 (Post-order)
        leftDeep := postOrder(node.Left)
        rightDeep := postOrder(node.Right)

        // 3. 현재 노드를 꺾임점으로 하는 지름으로 maxDiameter 갱신
        maxDiameter = max(maxDiameter, leftDeep + rightDeep)

        // 4. 부모 노드에게 제공할 "현재 노드의 최대 깊이" 리턴
        return max(leftDeep, rightDeep) + 1
    }

    postOrder(root)
    return maxDiameter
}