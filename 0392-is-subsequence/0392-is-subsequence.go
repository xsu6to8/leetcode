func isSubsequence(s string, t string) bool {
    if len(s) == 0 {
        return true
    }

    if len(s) > len(t) {
        return false
    }

    sIdx := 0
    for _, cmp := range t {
        target := s[sIdx]
        if byte(cmp) == target {
            sIdx++
            if sIdx == len(s) {
                return true
            }
        }
    }

    return false
}