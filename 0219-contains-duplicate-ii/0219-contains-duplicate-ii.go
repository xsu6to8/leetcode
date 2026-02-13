func containsNearbyDuplicate(nums []int, k int) bool {
    //  map <- key : val / value : idx
    nMap := make(map[int] int)

    for i := 0; i < len(nums); i++ {
        // 1. 처음 나온 수
        if val, exist := nMap[nums[i]]; !exist {
            nMap[nums[i]] = i
        } else {    // 2. 나온 적 있음 
            if i - val <= k {
                return true
            }
            nMap[nums[i]] = i   // 3. 나왔지만 k보다 큰 차이 -> 후에 나온 idx로 교체
        }
    }

    return false
}