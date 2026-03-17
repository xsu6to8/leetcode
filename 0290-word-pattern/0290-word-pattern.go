func wordPattern(pattern string, s string) bool {
    splitS := strings.Split(s, " ")
    if len(pattern) != len(splitS) {
        return false
    }

    charToWord := make(map[rune]string)
    wordToChar := make(map[string]rune)

    for i, char := range pattern {
        word := splitS[i]

        // 1. 패턴 -> 단어 체크
        if val, ok := charToWord[char]; ok {
            if val != word {
                return false
            }
        } else {
            // 2. 단어 -> 패턴 체크 (이미 다른 패턴에 할당된 단어인지)
            if _, exists := wordToChar[word]; exists {
                return false
            }
            charToWord[char] = word
            wordToChar[word] = char
        }
    }

    return true
}