/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    if headA == nil || headB == nil {
        return nil
    }

    destA, lenA := traverse(headA)
    destB, lenB := traverse(headB)
    
    // 끝까지 도달한 게 서로 다름 == intersect X
    if destA != destB {
        return nil
    }

    // 끝까지 도달한 게 같음 == intersect O
    // 역순으로 돌며 달라지는 지점 찾기
    lenDiff := lenA - lenB

    // B가 더 긴 경우
    if lenDiff < 0 {
        for ; lenDiff < 0; lenDiff++ {
            headB = headB.Next
        } 
    } else if lenDiff > 0 {
        for ; lenDiff > 0; lenDiff-- {
            headA = headA.Next
        } 
    }

    return findIntersect(headA, headB)
}

func traverse(node *ListNode) (*ListNode, int) {
    length := 1
    for {
        if node.Next == nil{
            break
        }
        node = node.Next 
        length++
    }
    return node, length
}

func findIntersect (n1, n2 *ListNode) *ListNode {
    for {
        if n1 == n2 {
            return n1
        }
        n1 = n1.Next
        n2 = n2.Next
    }
}