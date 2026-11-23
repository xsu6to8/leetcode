func minOperations(boxes string) []int {
    n := len(boxes)
    res := make([]int, n)

    total := 0 // 전체 공 개수
    for i := 0; i < n; i++ {
        if boxes[i] == '1' {
            // total 갱신
            total++
            // res[0]에 이 공의 기여분 더하기
            res[0] += i 
        }
    }

    left := 0 // 지금 목표 상자 "이하"에 있는 공 개수
    for i := 0; i < n-1; i++ {
        // i번 상자에 공이 있으면 left 갱신
        if boxes[i] == '1' {
            left++
        }

        right := total - left

        res[i+1] = res[i] + left - right
    }

    return res
}