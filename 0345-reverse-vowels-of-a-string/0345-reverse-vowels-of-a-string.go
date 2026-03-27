func reverseVowels(s string) string {
    vMap := map[rune]bool{
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }

    res := []rune(s) 
    
    l, r := 0, len(res)-1
    for l < r {
        for l < r && vMap[res[l]] == false {
            l++
        }

        for l < r && vMap[res[r]] == false {
            r--
        } 
        
        res[l], res[r] = res[r], res[l]
        l++
        r--
    }
    
    return string(res)
}