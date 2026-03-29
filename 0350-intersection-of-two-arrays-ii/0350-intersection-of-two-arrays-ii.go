func intersect(nums1 []int, nums2 []int) []int {
    set := make(map[int]int)
    var res []int

    for _, v := range nums1 {
        if set[v] == 0 {
            set[v] = 1
        } else {
            set[v] = set[v] + 1
        }
    }

    for _, v := range nums2 {
        if set[v] != 0 {
            res = append(res, v)
            set[v] = set[v] - 1
            if set[v] == 0 {
                delete(set, v)
            }
        }
    }

    return res
}