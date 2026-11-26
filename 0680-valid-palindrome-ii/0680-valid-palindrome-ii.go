func validPalindrome(s string) bool {
    l , r := 0, len(s)-1
    for l < r {
        if s[l] == s[r] {
            l++
            r--
        } else {
            l2, r2 := l+1, r
            isPalin2 := true
            for l2 < r2 {
                if s[l2] != s[r2] {
                    isPalin2 = false
                }
                l2++
                r2--
            }
            l3, r3 := l, r-1
            isPalin3 := true
            for l3 < r3 {
                if s[l3] != s[r3] {
                    isPalin3 = false
                }
                l3++
                r3--
            }

            return isPalin2 || isPalin3
        }
    }

    return true
}