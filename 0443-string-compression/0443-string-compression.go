func compress(chars []byte) int {
    // 단일 문자 case
    if len(chars) == 1 {
        return 1
    }

    
    l, r := 0, 0
    wrtPtr := 0
    for l < len(chars) {
        // 우선 문자 작성 + wrtPtr 한 칸 증가
        chars[wrtPtr] = chars[l]
        wrtPtr++

        cnt := 1
        r = l + 1

        // init 문자와 뒷문자가 같은 case
        // [ r < len(chars) ] -> 마지막 문자 1개인 case용 
        for r < len(chars) && chars[l] == chars[r] {
            cnt++
            r++
        }

        if cnt != 1 {
            if cnt < 10 {
                chars[wrtPtr] = byte('0' + cnt)
                wrtPtr++
            } else {
                var tmp [10]byte
                tmpCnt := -1
                for cnt > 0 {
                    tmpCnt++
                    tmp[tmpCnt] = byte('0' + cnt%10)
                    cnt /= 10
                }

                for ; tmpCnt >= 0; tmpCnt-- {
                    chars[wrtPtr] = tmp[tmpCnt]
                    wrtPtr++
                }
            }
        }

        l = r
    }

    return wrtPtr
}