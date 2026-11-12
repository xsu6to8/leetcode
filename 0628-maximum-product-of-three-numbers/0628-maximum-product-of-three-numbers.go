func maximumProduct(nums []int) int {
    n := len(nums)
    if n == 3 {
        return nums[0] * nums[1] * nums[2]
    }

    slices.Sort(nums)
    
    // 0 제외 전부 음수인 case
    if nums[n-1] == 0 {
        return 0
    }
    
    // 기댓값 : 양수 Top3
    fir := nums[n-1] * nums[n-2] * nums[n-3]
    // 기댓값 : 음수(절댓값) Top2 * 양수 Top1
    sec := nums[0] * nums[1] * nums[n-1]

    return max(fir, sec)
}