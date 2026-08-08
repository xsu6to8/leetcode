func hammingDistance(x int, y int) int {
    diff := x ^ y
    toBin := strconv.FormatInt(int64(diff), 2)

    cnt := 0
    for _, v := range toBin {
        if v == '1' {
            cnt++
        }
    }

    return cnt
}