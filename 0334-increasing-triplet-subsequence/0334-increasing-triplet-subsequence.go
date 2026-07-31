func increasingTriplet(nums []int) bool {
    // fir : 우선 0번 인덱스 기준 시작
    fir := nums[0]
    // sec : 가장 큰 값으로 시작 (작은 수로 점차 교체)
    sec := math.MaxInt
    
    // 1번 인덱스로 탐색
    for i := 1; i < len(nums); i++ {
        curr := nums[i]
        // [MaxInt보다 작은 수로 '갱신된 이후 얘들' 중] 얘보다 큰 얘 만났다 -> 이미 정렬된 상태 3개
        if sec < curr {
            return true
        } 

        // 가장 작은 수로 추정되는 얘보다 curr이 크다 -> 2번째로 작은 수 갱신
        if fir < curr {
            sec = curr
        } 

        // 가장 작은 수 인 줄 알았던 얘보다 curr이 작다 -> [실제 순서 관계는 작은 얘가 뒤지만]
        // 순서 관계가 깨질 수도 있으니, 앞에 있던 얘를 '가정하며' 값만 갱신
        // 1 4 6 -> 0 4 6도 될 수 있도록 (2번 케이스 예시)
        if curr < fir {
            fir = curr
        }
    }

    return false
}