func checkPerfectNumber(num int) bool {
    // negative OR 0 case
    if num <= 1 {
        return false
    }

    sum := 1
    for i := 2; i * i <= num; i++ {
        if num % i == 0 {
            sum += i
            if i * i != num { 
                sum += num / i
            }
        }
    }
    if sum == num {
        return true
    } else {
        return false
    }
}