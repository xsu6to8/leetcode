// import (
// 	"fmt"
// 	"strings"
// )
// func detectCapitalUse(word string) bool {
//     //  edge case : only 1 letter
//     if len(word) == 1 {
//         return true
//     }
    
//     var allCap func(str string) bool
//     allCap = func(str string) bool {
//         caps := strings.ToUpper(str)
//         if caps == str {
//             return true
//         }
//         return false
//     }

//     var allSmall func(str string) bool
//     allSmall = func(str string) bool {
//         smalls := strings.ToLower(str)
//         if smalls == str {
//             return true
//         }
//         return false
//     }

//     //  case1 : All capital
//     if allCap(word) {
//         return true
//     }

//     //  case2 : All small
//     if allSmall(word) {
//         return true
//     }

//     //  case3 : only 1st cap
//     if unicode.IsUpper(rune(word[0])) {
//         remain := word[1:]
        
//         if allSmall(remain) {
//             return true
//         }
//     }

//     return false
// }


// Ver2 : counting '# of Captical lettes'
func detectCapitalUse(word string) bool {
    capCount := 0
    
    for i := 0; i < len(word); i++ {
        if word[i] >= 'A' && word[i] <= 'Z' {
            capCount++
        }
    }

    // case1 : (capCount == N)
    // case2 : (capCount == 0)
    // case3 :(capCount == 1 &% [1st letter])
    return capCount == len(word) || capCount == 0 || (capCount == 1 && word[0] >= 'A' && word[0] <= 'Z')
}