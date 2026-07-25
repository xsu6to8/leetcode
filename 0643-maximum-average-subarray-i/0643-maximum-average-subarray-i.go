// sliding window -> 특정 칸 수씩 옮겨가며 한번의 순회로 해결

func findMaxAverage(nums []int, k int) float64 {
    currentSum := 0

    // init Window
    for i := 0; i < k; i++ {
        currentSum += nums[i]
    }
    
    maxSum := currentSum

    for i := k; i < len(nums); i++ {
        // 새로 들어오는 값은 더하고, 윈도우에서 벗어나는 맨 앞의 값은 제거
        currentSum = currentSum + nums[i] - nums[i-k]
        
        // 최댓값 갱신
        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return float64(maxSum) / float64(k)
}