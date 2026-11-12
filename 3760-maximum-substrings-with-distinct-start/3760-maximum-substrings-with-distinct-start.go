func maxDistinct(s string) int {
    cnt := 0
    distinctLetters := make(map[rune]struct{})
    for _, r := range s {
        if _, ok := distinctLetters[r]; !ok {
            distinctLetters[r] = struct{}{}
            cnt++
        } 
    }

    return cnt
}