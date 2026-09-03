func arrayPairSum(nums []int) int {
    sum := 0

    sort.Ints(nums)
    for i := len(nums)-2; i >= 0; i -= 2 {
        sum += nums[i]
    }

    return sum
}