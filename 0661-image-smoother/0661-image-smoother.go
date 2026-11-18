func imageSmoother(img [][]int) [][]int {
    rows := len(img)
    cols := len(img[0])

    res := make([][]int, rows)
    for i := range res {
        res[i] = make([]int, cols)
    }

    var calcAround func(x, y int) int
    calcAround = func(x, y int) int {
        dx := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
        dy := [8]int{1, 1, 1, 0, 0, -1, -1, -1}

        sum := img[x][y]
        cnt := 1
        for i := 0; i < 8; i++ {
            nx := x + dx[i]
            ny := y + dy[i]

            if nx >= 0 && nx < rows && ny >=0 && ny < cols {
                sum += img[nx][ny]
                cnt++
            }
        }

        return sum / cnt
    }

    for i := 0; i < rows; i++ {
        for j := 0; j < cols; j++ {
            res[i][j] = calcAround(i, j)
        }
    }

    return res
}