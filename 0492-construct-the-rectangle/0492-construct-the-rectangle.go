func constructRectangle(area int) []int {
    sqrt := int(math.Sqrt(float64(area)))

    var res []int
    gap := math.MaxInt
    for i := 1; i <= sqrt; i++ {
        if area % i == 0 {
            v1 := area / i
            v2 := i

            maxV := max(v1, v2)
            minV := min(v1, v2)
            tmpGap := v1 - v2

            if tmpGap < gap {
                res = res[:0]
                res = append(res, maxV, minV)
            }
        }
    }

    return res
}