func productExceptSelf(nums []int) []int {
    ltor := make([]int, len(nums)+1)
    rtol := make([]int, len(nums)+1)

    // real values start with '1st idxed'
    ltor[0], rtol[len(nums)] = 1, 1

    // left to right
    for i := 0; i < len(nums); i++ {
        curr := ltor[i] * nums[i]
        ltor[i+1] = curr
    }

    // right to left
    for i := 0; i < len(nums); i++ {
        curr := rtol[len(nums) - i] * nums[len(nums) - 1 -i]
        rtol[len(nums) - i - 1] = curr
    }

    var res []int
    for i := 1; i <= len(nums); i++ {
        curr := ltor[i - 1] * rtol[i]
        res = append(res, curr)
    } 

    return res
}