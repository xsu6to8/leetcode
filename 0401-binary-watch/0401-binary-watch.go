func readBinaryWatch(turnedOn int) []string {
    var res []string
    
    for h := 0; h < 12; h++ {
        for m := 0; m < 60; m++ {
            // h와 m의 binary 개수 더하기
            // 시계의 수들은 전부 [2의 제곱이기에 1의 수와 동일]
            if bits.OnesCount(uint(h)) + bits.OnesCount(uint(m)) == turnedOn {
                // 시간 형식에 맞춰 포맷팅
                res = append(res, fmt.Sprintf("%d:%02d", h, m))
            }
        }
    }
    return res
}