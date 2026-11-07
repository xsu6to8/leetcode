func findLHS(nums []int) int {
    maxL := 0
    
    slices.Sort(nums)
    l, r := 0, 0

    for l < len(nums) {
        cnt := 0
        curr := nums[l]
        for l + 1 < len(nums) && curr == nums[l+1] {
           cnt++
           l++ 
        }
        r = l + 1
        if r < len(nums) {
            next := nums[r]
            if next == curr + 1 {
                for r + 1 < len(nums) && next == nums[r+1] {
                    cnt++
                    r++
                }
                if cnt +2 > maxL {
                    maxL = cnt+2
                }
            }
        } 
        l++
    }

    return maxL
}