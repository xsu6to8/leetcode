func isAnagram(s string, t string) bool {
    // 길이 & 구성 components 일치해야 함

    // 1. 길이 검사
    if len(s) != len(t) {
        return false
    }
    
    // 2. 구성 components 검사
    uniMap := make(map[rune] int)
    for _, r := range s {
        // r = rune(r)  -> range로 뽑아내면 rune으로 자동 반환
        uniMap[r]++
    }  

    for _, r := range t {
        if uniMap[r] == 0 {
            return false
        } else {
            uniMap[r]--
            if uniMap[r] == -1 {
                return false
            }
        }
    }  

    for _, val := range uniMap {
        if val != 0 {
            return false
        }
    }

    return true
}