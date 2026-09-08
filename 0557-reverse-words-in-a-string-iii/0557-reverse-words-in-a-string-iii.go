func reverseWords(s string) string {
    splited := strings.Fields(s)
    
    var res string
    for i := 0; i < len(splited); i++ {
        curr := splited[i]
        runes := []rune(curr)
        for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		    runes[i], runes[j] = runes[j], runes[i]
	    }
        res += string(runes) + " "
    }

    res = res[:len(res)-1]

    return res
}