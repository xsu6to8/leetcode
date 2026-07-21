func reverseVowels(s string) string {
    vowels := map[byte]bool{
        'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
        'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
    }

    //  Go에서 string은 직접 수정 불가 -> byte 슬라이스로 변환
    bs := []byte(s)

    l, r := 0, len(bs) - 1
    possibleL, possibleR := false, false
    
    for l < r { 
        if vowels[bs[l]] {
            possibleL = true
        } else {
            l++
        }

        if vowels[bs[r]] {
            possibleR = true
        } else {
            r--
        }

        // 둘 다 모음을 찾음 -> swap
        if possibleL && possibleR {
            tmp := bs[l]
            bs[l] = bs[r]
            bs[r] = tmp
            
            l++
            r--
            possibleL = false
            possibleR = false
        }
    } 

    return string(bs)
}