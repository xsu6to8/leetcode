/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS 요구 -> Queue
type Queue struct {
    items []*TreeNode
}

// Enqueue (뒤에 추가)
func (q *Queue) Push(val *TreeNode) {
    if val == nil {
        return
    }
    q.items = append(q.items, val)
}

// Dequeue (앞에서 제거)
func (q *Queue) Pop() (*TreeNode) {
    if len(q.items) == 0 {
        return nil
    }
    val := q.items[0]
    q.items = q.items[1:]
    return val
}

func averageOfLevels(root *TreeNode) []float64 {
    var res []float64
    
    q := new(Queue)
    q.Push(root)

    for len(q.items) != 0 {
        sum := 0
        qLen := len(q.items)
        for i := 0; i < qLen; i++ {
            currNode := q.Pop()
            sum += currNode.Val
            q.Push(currNode.Left)
            q.Push(currNode.Right)
        }

        res = append(res, float64(sum) / float64(qLen))
    }

    return res
}