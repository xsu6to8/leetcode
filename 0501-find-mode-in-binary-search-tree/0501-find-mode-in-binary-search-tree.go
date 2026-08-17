/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    if root == nil {
        return nil
    }
    if root.Left == nil && root.Right == nil {
        return []int{root.Val}
    }
    
    var res []int

    var prev *TreeNode
    cnt := 0 
    maxCnt := 0
    // in-order search -> BST 이기에 in-order 시, 크기순 정렬된 순서 탐방
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        // exit conditon
        if node == nil {
            return
        }

        inorder(node.Left)

    // main logic
        // 1. First vistied node
        if prev == nil {
            prev = node
            cnt = 1
        } else { // 2. non-First node
            if prev.Val == node.Val {
                cnt++
            } else if prev.Val != node.Val {
                if cnt == maxCnt{
                   res = append(res, prev.Val)     
                } else if cnt > maxCnt {
                    maxCnt = cnt
                    
                    // [res init] for more Bigger cnt
                    res = res[:0]
                    res = append(res, prev.Val)                  
                } 
                // new Val assiciation for 'cnt'
                cnt = 1
            }

            prev = node
        }

        inorder(node.Right)
    }

    inorder(root)

    // last node
    if cnt == maxCnt{
        res = append(res, prev.Val)     
    } else if cnt > maxCnt {
        maxCnt = cnt
                    
        // [res init] for more Bigger cnt
         res = res[:0]
         res = append(res, prev.Val)                  
    }  

    return res
}