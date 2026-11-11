// func isStrictlyPalindromic(n int) bool {
//     var baseConvert func(n, base int) bool
//     baseConvert = func(n, base int) bool {
//         converted := ""
//         for n > 0 {
//             curr := n % base
//             converted += strconv.Itoa(curr)

//             n /= base
//         }
//         l, r := 0, len(converted)-1
//         for l < r {
//             if converted[l] != converted[r] {
//                 return false
//             }
//             l++
//             r--
//         }
//         return true
//     }

//     for i := 2; i < n-1; i++ {
//         if !baseConvert(n, i) {
//             return false
//         }
//     }

//     return true
// }

func isStrictlyPalindromic(n int) bool {
    // n-2 진법이면 항상 "12"를 만들기에 무조건 false
    return false
}