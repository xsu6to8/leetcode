/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    switch helper(root){
        case -1 :
            return false
        default : 
            return true
    }
}

func helper(node *TreeNode) int {
    // 1. Base Case
    if node == nil {
        return 0
    }

    // 2. 왼쪽 서브트리 높이
    left := helper(node.Left)
    // 만약 왼쪽에서 이미 균형이 깨졌다면 -> [-1]
    if left == -1{
        return -1
    }

    // 3. 오른쪽 서브트리의 높이
    right := helper(node.Right)
    // 만약 오른쪽에서 이미 균형이 깨졌다면 -> [-1] 
    if right == -1{
        return -1
    }

    // 4. 왼쪽 높이와 오른쪽 높이의 차이가 1보다 큼 -> [-1]
    if Abs(left - right) > 1{
        return -1
    }

    // 5. 균형이 잡혀있다면, 현재 노드 높이를 계산 -> 부모에 전달
    currH := max(left, right) + 1
    return currH 
}

func Abs(n int) int {
    if n < 0 {
        return -n
    }
    return n
}