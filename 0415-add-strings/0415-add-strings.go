func addStrings(num1 string, num2 string) string {
    i := len(num1) - 1
    j := len(num2) - 1

    var res string

    isCarry := false
    for i >= 0 || j >= 0 || isCarry {
        currSum := 0
        if i >= 0 {
            currSum += (int)(num1[i] - '0')
            i--
        }
        if j >= 0 {
            currSum += (int)(num2[j] - '0')
            j--
        }

        if isCarry {
            currSum++
            isCarry = false
        }

        if currSum >= 10 {
            isCarry = true
            currSum %= 10
        }

        res =string(byte(currSum + '0')) + res
    }

    return res
}