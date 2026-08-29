func nextPermutation(nums []int)  {
    var isChanged bool
    for i := len(nums)-1; i > 0; i-- {
        if nums[i-1] < nums[i] {
            isChanged = true
            idx := i-1
            for j := len(nums)-1; j > idx; j-- {
                if nums[j] > nums[idx] {
                    nums[j], nums[idx] = nums[idx], nums[j]
                    break
                }
            }
            for k, l := idx+1, len(nums)-1; k < l; k, l = k+1, l-1 {
                nums[k], nums[l] = nums[l], nums[k]
            }
            break
        } 
    }
    if !isChanged {
        for k, l := 0, len(nums)-1; k < l; k, l = k+1, l-1 {
            nums[k], nums[l] = nums[l], nums[k]
        }
    }
}