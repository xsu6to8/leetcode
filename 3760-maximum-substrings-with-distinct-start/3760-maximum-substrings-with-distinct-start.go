// func maxDistinct(s string) int {
//     distinctLetters := make(map[rune]struct{})
//     for _, r := range s {
//         if _, ok := distinctLetters[r]; !ok {
//             distinctLetters[r] = struct{}{}
//         } 
//     }

//     return len(distinctLetters)
// }

func maxDistinct(s string) int {
    var mask uint32 

    for _, r := range s {
        // r에 해당하는 비트를 mask에 켜기
        shiftCnt := r - 'a'
        curr := uint32(1) << shiftCnt
        mask |= curr

        // 알파벳 26개가 모두 나오면 조기 종료
        if bits.OnesCount32(mask) == 26 {
            return 26
        }
    }

    // 켜진 비트 개수 = 서로 다른 문자 수
    return bits.OnesCount32(mask)
}