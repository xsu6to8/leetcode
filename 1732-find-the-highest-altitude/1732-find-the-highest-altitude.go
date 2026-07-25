func largestAltitude(gain []int) int {
    var res int
    if gain[0] < 0 {
        res = 0
    } else {
        res = gain[0]
    }

    for i := 1; i < len(gain); i++ {
        gain[i] = gain[i] + gain[i-1]
        if gain[i] > res {
            res = gain[i]
        }
    }

    return res
}