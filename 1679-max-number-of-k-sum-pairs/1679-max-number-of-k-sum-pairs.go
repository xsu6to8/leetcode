func maxOperations(nums []int, k int) int {
    sort.Ints(nums)
    
    cnt := 0
    l, r := 0, len(nums) - 1
    for l < r {
        curr := nums[l] + nums[r]
        if curr < k {
            l++
        } else if curr > k {
            r--
        } else {
            cnt++
            l++
            r--
        }
    }

    return cnt
}