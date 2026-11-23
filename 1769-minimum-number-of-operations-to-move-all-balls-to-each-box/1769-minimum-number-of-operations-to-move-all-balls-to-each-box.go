func minOperations(boxes string) []int {
    var idxSlice []int
    for i, v := range boxes {
        if v == '0' {
            continue
        } else {
            idxSlice = append(idxSlice, i)
        }
    }

    var abs func(a, b int) int 
    abs = func(a, b int) int{
        if b > a {
            return b - a
        } else {
            return a - b
        }
    }

    var res []int
    for i, _ := range boxes {
        sum := 0
        for _, j := range idxSlice {
            sum += abs(i, j)
        }
        res = append(res, sum)
    }

    return res
}