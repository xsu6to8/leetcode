func isHappy(n int) bool {
    appear := make(map[int] bool)

    for {
        if n == 1 {
            return true
        } else {
            _, exist := appear[n]
            if exist == true {
                return false
            }

            appear[n] = true
            n = getNext(n)
        }
    }
}

func getNext(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}