func isPowerOfTwo(n int) bool {
    if n <= 0 {
        return false
    }

    if n == 1 {
        return true
    }
    
    for n != 2 {
        if n & 0x01 == 1 {
            return false
        } else {
            n >>= 1
        }
    }

    return true
}