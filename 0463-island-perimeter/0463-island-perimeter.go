func islandPerimeter(grid [][]int) int {
    dx := []int{-1, 1, 0, 0}
    dy := []int{0, 0, -1, 1}

    res := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == 0 {
                continue
            } else {
                res += 4
                for k := 0; k < 4; k++ {
                    nx := i + dx[k]
                    ny := j + dy[k]

                    if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[i]) {
                        if grid[nx][ny] == 1{
                            res--
                        }
                    }
                } 
            }
        }
    }

    return res
}