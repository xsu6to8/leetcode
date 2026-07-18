// func removeStars(s string) string {
//     var resStack []rune

//     for _, v := range(s) {
//         if v != '*' {
//             resStack = append(resStack, v)
//         } else {
//             top := len(resStack) - 1
// 		    resStack = (resStack)[:top]   // 스택에 마지막 데이터 제거함
//         }
//     }

//     res := string(resStack)
//     return res
// }


// 성능 개선 시, 투 포인터 가능
func removeStars(s string) string {
    b := []byte(s)
    
    // 값을 새로 쓸 위치를 가리키는 포인터 (스택의 top 역할)
    writeIdx := 0
    
    for i := 0; i < len(b); i++ {
        if b[i] == '*' {
            // 별을 만나면 쓰기 포인터를 하나 다시 앞으로 (정상 진행 방향 ->)
            if writeIdx > 0 {
                writeIdx--
            }
        } else {
            // 일반 문자면 현재 쓰기 위치에 값을 덮어쓰고 포인터 뒤로 (반대 진행 방향 <-)
            b[writeIdx] = b[i]
            writeIdx++
        }
    }

    return string(b[:writeIdx])
}