/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
    // 0개 또는 1개의 노드가 있을 때
    if head == nil || head.Next == nil {
        return head
    }
    
    odd := head
    even := head.Next
    evenHead := even
    for even != nil && even.Next != nil{
        odd.Next = even.Next
        odd = odd.Next
        even.Next = odd.Next
        even = even.Next
    }

    odd.Next = evenHead

    return head
}