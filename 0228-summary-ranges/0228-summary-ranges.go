func summaryRanges(nums []int) []string {
    res := []string{}

    for i := 0; i < len(nums); i++ {
        var tmp string
        tmp = strconv.Itoa(nums[i])

        incr := false
        for {
            if  i < len(nums) - 1 && nums[i+1] - nums[i] == 1 {
                i++
                incr = true
            } else {
                break
            }
        }

        if incr {
            last := strconv.Itoa(nums[i])
            tmp = tmp + "->" + last
        }

        res = append(res, tmp)
    }

    return res
}