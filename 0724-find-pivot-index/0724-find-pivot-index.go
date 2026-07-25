func pivotIndex(nums []int) int {
    total := 0
    for _, num := range nums {
        total += num
    }
    total -= nums[0]
    
    // idx : 0
    if total == 0 {
        return 0
    }

    sum := 0
    for i := 1; i < len(nums); i++ {
        sum += nums[i-1]
        total -= nums[i]

        if sum == total {
            return i
        }
    }

    return -1
}