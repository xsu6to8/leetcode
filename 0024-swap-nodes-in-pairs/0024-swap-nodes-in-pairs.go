/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
    dummy := &ListNode{Next : head}
    prev := dummy
    
    for prev.Next != nil && prev.Next.Next != nil {
        // init
        fir := prev.Next
        sec := prev.Next.Next

        // swap
        fir.Next = sec.Next
        sec.Next = fir
        prev.Next = sec
        
        // mve to next pair
        prev = fir
    }

    return dummy.Next
}