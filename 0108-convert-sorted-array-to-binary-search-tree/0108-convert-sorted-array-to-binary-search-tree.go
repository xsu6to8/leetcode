/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left int, right int) *TreeNode {
    // 탈출 조건
    if left > right {
        return nil
    }
    
    // 중간 인덱스
    mid := (left + right) / 2

    // 현재 mid 값 노드 생성
    node := new(TreeNode)
    node.Val = nums[mid]
    
    // 재귀
    node.Left = helper(nums, left, mid - 1)
    node.Right = helper(nums, mid + 1, right)

    // 생성 노드 반환
    return node
}