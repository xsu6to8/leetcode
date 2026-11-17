func judgeCircle(moves string) bool {
    ud := 0
    lr := 0
    for _, c := range moves {
        if c == 'U' {
            ud++
        }
        if c == 'D' {
            ud--
        }

        if c == 'L' {
            lr--
        }
        if c == 'R' {
            lr++
        }
    }

    if ud == 0 && lr == 0 {
        return true
    }

    return false
}