/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    slow, fast := head, head
    // fast와 slow의 차이 : n 만큼 벌어짐
    for i := 0; i < n; i++ {
        fast = fast.Next
    }

    // n : 마지막 노드 지우기
    if fast == nil {
        return head.Next
    }

    // fast 끝까지 보내기
    for fast.Next != nil {
        slow = slow.Next
        fast = fast.Next
    }

    slow.Next = (slow.Next).Next

    return head
}