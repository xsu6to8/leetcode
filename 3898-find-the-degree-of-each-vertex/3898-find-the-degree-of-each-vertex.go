func findDegrees(matrix [][]int) []int {
    var res []int
    for _, sli := range matrix {
        cnt := 0
        for _, v := range sli {
            if v == 1 {
                cnt++
            }
        }
        res = append(res, cnt) 
    }
    return res
}