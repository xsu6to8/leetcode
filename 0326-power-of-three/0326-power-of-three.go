func isPowerOfThree(n int) bool {
    // 가용 범위 내 가장 큰 3의 제곱수
    return n > 0 && 1162261467 % n == 0
}