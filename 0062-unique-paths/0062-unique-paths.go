func uniquePaths(m int, n int) int {
    // Moving > width: m-1, height: n-1
    total := m + n - 2
    
    // 'Combination' minimization
    k := m - 1
    if m - 1 > n - 1 {
        k = n - 1
    }
    
    // nCk
    res := 1
    for i := 1; i <= k; i++ {
        res = res * (total - i + 1) / i
    }
    
    return res
}