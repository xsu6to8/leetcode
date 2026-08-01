func thirdMax(nums []int) int {
    removedDups := []int{}
    iMap := make(map[int]struct{})
    for _, v := range nums {
        if _, ok := iMap[v]; !ok {
            iMap[v] = struct{}{}
            removedDups = append(removedDups, v)
        } 
    }

    // third 자체가 없는 case
    if len(removedDups) < 3 {
        return slices.Max(removedDups)
    }

    fir := math.MinInt 
    sec := math.MinInt 
    thi := math.MinInt 
    for _, v := range removedDups {
        if v > fir {
            tmpFir := fir
            fir = v
            tmpSec := sec
            sec = tmpFir
            thi = tmpSec
        } else if v < fir && v > sec {
            tmpSec := sec
            sec = v
            thi = tmpSec
        } else if v < sec && v > thi {
            thi = v
        }     
    }

    return thi
}