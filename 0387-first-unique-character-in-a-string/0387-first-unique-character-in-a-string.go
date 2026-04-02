func firstUniqChar(s string) int {
    countMap := make(map[rune]int)
    for _, char := range s {
        countMap[char]++
    }

    // map : iteration 순서 보장 X
    for i, char := range s {
        if countMap[char] == 1 {
            return i
        }
    }

    return -1
}

// 1. 지금처럼 알파벳 only일 땐 map보다 [26]int 배열이 유리 (해시 연산/메모리 할당 오버헤드 제거)
// 2. 배열은 연속된 메모리 구조로 cache hit가 높아 map보다 실제 실행 속도가 훨씬 빠름