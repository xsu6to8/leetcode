// func checkRecord(s string) bool {
//     aCnt := 0

//     for i := 0; i < len(s); i++ {
//         if s[i] == 'A' {
//             aCnt++
//             if aCnt == 2 {
//                 return false
//             }
//         } else if s[i] == 'L' {
//             if i+1 < len(s) && i+2 < len(s) && s[i+1] == 'L' && s[i+2] == 'L' {
//                 return false
//             }
//         }
//     }

//     return true
// }

// 전부 내장함수로 해결하는 version
func checkRecord(s string) bool {
    return strings.Count(s, "A") < 2 && !strings.Contains(s, "LLL")
}