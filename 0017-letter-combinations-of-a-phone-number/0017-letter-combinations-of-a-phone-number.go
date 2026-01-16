func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    mapping := []string{
        "", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
    }
    
    var res []string

    var backtrack func(index int, currentStr string)
    backtrack = func(index int, currentStr string) {
        if len(currentStr) == len(digits) {
            res = append(res, currentStr)
            return
        }

        digit := digits[index] - '0'
        letters := mapping[digit]

        for _, char := range letters {
            backtrack(index + 1, currentStr + string(char))
        }
    }

    backtrack(0, "")
    
    return res
}