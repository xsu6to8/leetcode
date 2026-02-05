func isAlphanumeric(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func toLower(c byte) byte {
    if c >= 'A' && c <= 'Z' {
        return c + ('a' - 'A')
    }
    return c
}

func isPalindrome(s string) bool {
    l := 0
    r := len(s) - 1

    for l < r {
        if !isAlphanumeric(s[l]) {
            l++
            continue
        }

        if !isAlphanumeric(s[r]) {
            r--
            continue
        }

        if toLower(s[l]) != toLower(s[r]) {
            return false
        }

        l++
        r--
    }

    return true
}
