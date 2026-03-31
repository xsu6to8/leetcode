func intToRoman(num int) string {
    values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
    symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
    
    res := ""
    idx := 0
    for idx < 13 {
        if num >= values[idx] {
            num -= values[idx]
            res += symbols[idx]
        } else {
            idx += 1
        }
    }

    return res
}