type NumArray struct {
    arr []int
}


func Constructor(nums []int) NumArray {
    na := &NumArray{nums}
    return *na
}


func (this *NumArray) SumRange(left int, right int) int {
    sum := 0
    for left <= right {
        sum += this.arr[left]
        left++
    }

    return sum
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */