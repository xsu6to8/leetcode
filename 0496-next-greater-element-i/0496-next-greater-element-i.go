func nextGreaterElement(nums1 []int, nums2 []int) []int {
    // Build nums2 idx map : O(nums2.length)
    var stack []int
    num2Map := make(map[int]int)
    for i := 0; i < len(nums2); i++ {
        curr := nums2[i]

        // non-empty stack
        for len(stack) > 0 && stack[len(stack)-1] < curr {
            top := stack[len(stack)-1]
            num2Map[top] = curr
            stack = stack[:len(stack)-1] // pop
        }

        stack = append(stack, curr)
    }

    // remainders in  stack
    for _, v := range stack {
        num2Map[v] = -1
    }

    // loop on nums1 : O(nums1.length)
    var res []int
    for _, v := range nums1 {
        res = append(res, num2Map[v])
    }     

    return res 
}