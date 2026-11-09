/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
    var gcd func(a, b int) int
    gcd = func(a, b int) int {
        if b == 0 {
            return a
        }
        return gcd(b, a % b)
    }
    
    curr := head
    for curr.Next != nil {
        tmp := curr

        cVal := curr.Val
        nVal := curr.Next.Val

        gcdVal := gcd(cVal, nVal)
        newNode := new(ListNode)
        newNode.Val = gcdVal
        newNode.Next = tmp.Next
        tmp.Next = newNode

        curr = curr.Next.Next
    } 

    return head
}