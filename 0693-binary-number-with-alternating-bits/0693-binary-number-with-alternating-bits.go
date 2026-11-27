func hasAlternatingBits(n int) bool {
    cmp := n ^ (n >> 1)
    if cmp & (cmp+1) == 0 {
        return true
    }

    return false
}