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

/*
이 문제 국한하자면 array [26]int로 가고
ransom은 + / magazine은 - 역할만 수행하며
중간에 negative 값 띄면 바로 false
마지막에 다 0인지 확인하는 게 더 효율적

만일 근데 utf-8 모든 유니코드 수용이라면 map 방식이 좋음
*/