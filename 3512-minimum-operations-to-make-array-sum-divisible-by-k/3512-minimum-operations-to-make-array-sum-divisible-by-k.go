func minOperations(nums []int, k int) int {
    sum := 0
    for _, v := range nums {
        sum += v
    }
    if sum < k {
        return sum
    }
    return sum % k
}