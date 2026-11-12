func maxDistinct(s string) int {
    distinctLetters := make(map[rune]struct{})
    for _, r := range s {
        if _, ok := distinctLetters[r]; !ok {
            distinctLetters[r] = struct{}{}
        } 
    }

    return len(distinctLetters)
}