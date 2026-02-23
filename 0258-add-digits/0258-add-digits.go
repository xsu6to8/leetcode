func addDigits(num int) int {
    var res int
    for {   //  escape condition
        sum := 0
        for num != 0 {   //  inner logic
            sum += num % 10
            num /= 10 
        }

        if sum < 10 {
            res = sum
            break
        }

        num = sum
    }

    return res
}