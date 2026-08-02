func arrangeCoins(n int) int {
    for i := 1;  ; i++ {
        if 2 * n <(i * (i+1)) {
            return i -1
        }
    }

    return 0
}