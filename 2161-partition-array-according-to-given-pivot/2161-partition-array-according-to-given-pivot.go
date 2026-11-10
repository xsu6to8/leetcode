func pivotArray(nums []int, pivot int) []int {
    var small []int
    var same []int
    var big []int
    for _, v := range nums {
        if v < pivot {
            small = append(small, v)
        } else if v == pivot {
            same = append(same, v)
        } else {
            big = append(big, v)
        }
    }

    res := append(small, same...)
    res = append(res, big...)
    return res
}