func isPowerOfFour(n int) bool {
    // 2의 제곱수
    if n > 0 && (n & (n - 1)) == 0 {
        // 4의 제곱수
        if n % 3 == 1 {
            return true
        } else if n % 3 == 2 { // 2의 제곱수
            return false
        }
    }
    return false
}