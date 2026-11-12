// 사실 상, 1 개수 맘대로 쓸 수 있다는 조건 -> 가장 큰 수 찾기
func minPartitions(n string) int {
    sTob := []byte(n)
    maxN := 0
    for _, v := range sTob {
        val := int(v - '0')
        if val > maxN {
            maxN = val
        }
    }

    return maxN
}