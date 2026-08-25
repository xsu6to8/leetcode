/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

    /*
        [1,null,3,2]
        와 같이 왼쪽은 없지만 오른쪽 자식만 있는 경우에는

        //  BST -> in-order search == search in size order

func getMinimumDifference(root *TreeNode) int {

    minGap := math.MaxInt



    var prev *TreeNode

    var inOrder func(node *TreeNode)

    inOrder = func(node *TreeNode) {

        if node == nil {

            return

        }

       

        inOrder(node.Left)

       

        // case1 :  most-deepest & leftest node

        if prev == nil {

            prev = node

            return

        }

        // case2 :  main logic

        //          = [current value - previous value]

        currGap := node.Val - prev.Val

        if currGap < minGap {

            minGap = currGap

        }



        prev = node



        inOrder(node.Right)

    }



    inOrder(root)

    return minGap

}
    logic 깨짐 <- 중간의 case1 에서
    */

// 중위순회로 slicde에 값 저장하고 비교
func getMinimumDifference(root *TreeNode) int {
    var vals []int
    var inOrder func(node *TreeNode)
    inOrder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inOrder(node.Left)
        vals = append(vals, node.Val)
        inOrder(node.Right)
    }
    
    inOrder(root)
    
    minGap := math.MaxInt
    for i := 1; i < len(vals); i++ {
        if vals[i] - vals[i-1] < minGap {
            minGap = vals[i] - vals[i-1]
        }
    }
    return minGap
}