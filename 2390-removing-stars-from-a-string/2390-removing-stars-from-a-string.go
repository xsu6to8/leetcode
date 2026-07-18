func removeStars(s string) string {
    var resStack []rune

    for _, v := range(s) {
        if v != '*' {
            resStack = append(resStack, v)
        } else {
            top := len(resStack) - 1
		    resStack = (resStack)[:top]   // 스택에 마지막 데이터 제거함
        }
    }

    res := string(resStack)
    return res
}