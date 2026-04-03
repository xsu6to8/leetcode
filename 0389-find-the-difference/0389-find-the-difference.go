func findTheDifference(s string, t string) byte {
    var sSlice [26]rune
    for _, c := range s {
        sSlice[c- 'a']++
    }
    for _, c := range t {
        sSlice[c - 'a']--
    }
    for i, v := range sSlice {
        if v == -1 {
            return byte('a' + i)
        }
    }

    return 'a'
}