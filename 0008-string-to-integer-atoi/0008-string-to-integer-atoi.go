func myAtoi(s string) int {
    s = strings.TrimSpace(s)
    if len(s) == 0 {
        return 0
    }

    res := 0
    sign := 1
    i := 0

    if s[i] == '-' {
        sign = -1
        i++
    } else if s[i] == '+' {
        i++
    }

    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        digit := int(s[i] - '0')

        if res > (2147483647-digit)/10 {
            if sign == 1 {
                return 2147483647
            }
            return -2147483648
        }

        res = res*10 + digit
        i++
    }

    return res * sign
}