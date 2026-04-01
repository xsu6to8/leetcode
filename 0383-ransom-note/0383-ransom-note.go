func canConstruct(ransomNote string, magazine string) bool {
    rMap := make(map[rune]int)
    for _, c := range ransomNote {
        rMap[c]++
    }

    for _, c := range magazine {
        if rMap[c] == 0 {
            continue
        } else {
            rMap[c]--
        }
    }

    for _, v := range rMap {
        if v != 0 {
            return false
        }
    }

    return true
}