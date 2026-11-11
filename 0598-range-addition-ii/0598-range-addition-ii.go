func maxCount(m int, n int, ops [][]int) int {
    if len(ops) == 0 {
        return m*n
    }

    minX, minY := m, n
    for _, v := range ops {
        currX, currY := v[0], v[1]
        if currX < minX {
            minX = currX
        }
        if currY < minY {
            minY = currY
        }
    }

    return minX*minY
}