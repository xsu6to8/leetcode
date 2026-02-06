/**
 * Definition for singly-linked list.
 * type ListNode struct {
 * Val int
 * Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return false
    }

    oneStep := head
    twoStep := head

    for {
        if twoStep == nil || twoStep.Next == nil {
            return false
        }

        oneStep = oneStep.Next
        twoStep = twoStep.Next.Next

        if oneStep == twoStep {
            return true
        }
    }
}