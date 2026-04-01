func threeSumClosest(nums []int, target int) int {
    slices.Sort(nums)
    n := len(nums)
    closestSum := nums[0] + nums[1] + nums[2]

    for i := 0; i < n-2; i++ {
        left, right := i+1, n-1
        
        for left < right {
            currentSum := nums[i] + nums[left] + nums[right]
            
            if abs(target - currentSum) < abs(target - closestSum) {
                closestSum = currentSum
            }

            if currentSum < target {
                left++ 
            } else if currentSum > target {
                right-- 
            } else {
                return currentSum 
            }
        }
    }
    return closestSum
}

func abs(x int) int {
    if x < 0 { 
        return -x 
    }
    return x
}