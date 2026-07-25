func findMaxAverage(nums []int, k int) float64 {
    maxAvg := -math.MaxFloat64

    for i := 0; i <= len(nums) - k; i++ {
        sum := 0
        for j := 0; j < k; j++ {
            sum += nums[i+j]
        } 

        currAvg := float64(sum) / float64(k)
        if currAvg > maxAvg {
            maxAvg = currAvg
        }
    }

    return maxAvg
}