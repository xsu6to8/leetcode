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