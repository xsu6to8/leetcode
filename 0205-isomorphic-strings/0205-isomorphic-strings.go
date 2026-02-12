func isIsomorphic(s string, t string) bool {
    StoT := make(map[byte] byte)
    TtoS := make(map[byte] byte)

    for i := 0; i < len(s); i++ {
        if varT, exist := StoT[s[i]]; !exist {
            StoT[s[i]] = t[i]
        } else {
            // map에 존재하는 경우
            // 1. 이미 다른 문자라면
            if varT != t[i]{
                return false
            } else { // 2. t[i] 맞다면 통과

            }
        }

        if varS, exist := TtoS[t[i]]; !exist {
            TtoS[t[i]] = s[i]
        } else {
            // map에 존재하는 경우
            // 1. 이미 다른 문자라면
            if varS != s[i]{
                return false
            } else { // 2. t[i] 맞다면 통과

            }
        }
    }

    return true
}