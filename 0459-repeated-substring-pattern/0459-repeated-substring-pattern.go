func repeatedSubstringPattern(s string) bool {
    testStr := s + s
    testStr = testStr[1:len(testStr)-1]

    originLen := len(s)
    testLen := len(testStr)

    isMatched := false
    for i := 0; i <= testLen - originLen; i++ {
        if s == testStr[i : i + originLen] {
            isMatched = true
        }
    }

    return isMatched
}