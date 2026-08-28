func divide(dividend int, divisor int) int {
    // edge case: −2^31 / -1 
    if dividend == math.MinInt32 && divisor == -1 {
        return math.MaxInt32
    }

    // Neg가 더 넓은 범위이기에 음수로 바꾸고 시작 (−2^31 을 다루는 경우를 위해)
    var toNeg func(n int) int
    toNeg = func(n int) int {
        if n > 0 {
            return -n
        }
        return n
    }
    
    isPos := (dividend > 0 && divisor > 0) || (dividend < 0 && divisor < 0)

    dividend = toNeg(dividend)
    divisor = toNeg(divisor)
    quotient := 0
    for {
        if dividend > divisor { 
            break 
        }
        curDivisor := divisor
        curQuotient := -1 

        // 2배로 키워도 dividend보다 절댓값이 커지지 않는 동안
        // by Overflow 방지 체크
        for curDivisor >= (math.MinInt32 >> 1) && dividend <= (curDivisor << 1) { 
            curDivisor <<= 1
            curQuotient <<= 1
        }

        dividend -= curDivisor
        quotient += curQuotient
    }

    if isPos {
        return -quotient
    } else {
        return quotient
    }
}