// func repeatedSubstringPattern(s string) bool {
//     testStr := s + s
//     testStr = testStr[1:len(testStr)-1]

//     originLen := len(s)
//     testLen := len(testStr)

//     isMatched := false
//     // maximum : O(2N)
//     for i := 0; i <= testLen - originLen; i++ {
//         // inner logic -> O(N) 일일히 비교
//         if s == testStr[i : i + originLen] {
//             isMatched = true
//         }
//     }

//     return isMatched
// }


func repeatedSubstringPattern(s string) bool {
    // s + s 생성 후 양 끝 1글자씩 제거
    // -> 같은 문자열 존재 시, [one cycle 돌리면 한번은 나오게 되어있음]
    testStr := (s + s)[1 : len(s)*2-1]
    
    // strings.Contains 함수 -> 부분 문자열 포함 여부 검사 O(N)
    return strings.Contains(testStr, s)
}