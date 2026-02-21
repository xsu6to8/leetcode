/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    slow, fast := head, head
    // Empty Linked List case
    if slow == nil {
        return false
    }

    // fast가 끝에 도달할 때까지 slow는 중간으로 이동
    for fast != nil && fast.Next != nil {
        slow = slow.Next
        fast = fast.Next.Next
    }

    var prev *ListNode
    curr := slow

    for curr != nil {
        nextTemp := curr.Next 
        curr.Next = prev      
        prev = curr           
        curr = nextTemp       
    }

    left, right := head, prev
    for right != nil { // 뒤집힌 리스트(길이가 같거나 1 작음) 기준으로 비교
        if left.Val != right.Val {
            return false
        }
        left = left.Next
        right = right.Next
    }
    return true
}