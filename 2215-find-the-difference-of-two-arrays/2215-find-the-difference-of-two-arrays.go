func findDifference(nums1 []int, nums2 []int) [][]int {
    map1 := make(map[int]bool)
    for _, v := range nums1 {
        map1[v] = true
    }
    map2 := make(map[int]bool)
    for _, v := range nums2 {
        map2[v] = true
    }

    res := make([][]int, 2)
    for v, _ := range map1 {
        if map2[v] == true {
            continue
        }
        res[0] = append(res[0], v)
    }
    for v, _ := range map2 {
        if map1[v] == true {
            continue
        }
        res[1] = append(res[1], v)
    }

    return res
}