// type NumArray struct {
//     arr []int
// }


// func Constructor(nums []int) NumArray {
//     na := &NumArray{nums}
//     return *na
// }


// func (this *NumArray) SumRange(left int, right int) int {
//     sum := 0
//     for left <= right {
//         sum += this.arr[left]
//         left++
//     }

//     return sum
// }


// /**
//  * Your NumArray object will be instantiated and called as such:
//  * obj := Constructor(nums);
//  * param_1 := obj.SumRange(left,right);
//  */


// 최적화 - 누적합 => {O(n) -> O(1)}
type NumArray struct {
    sums []int
}

func Constructor(nums []int) NumArray {
    n := len(nums)
    // 계산 편의를 위해 길이를 n+1로 설정 (0번째 인덱스는 0)
    sums := make([]int, n+1)
    for i := 0; i < n; i++ {
        sums[i+1] = sums[i] + nums[i]
    }
    return NumArray{sums: sums}
}

func (this *NumArray) SumRange(left int, right int) int {
    return this.sums[right+1] - this.sums[left]
}