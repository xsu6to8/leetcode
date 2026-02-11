func titleToNumber(columnTitle string) int {
    res := 0
    
    for i := 0; i < len(columnTitle); i++ {
        tmp := columnTitle[i] - 'A' + 1
        res = res*26 + int(tmp)
    }
    
    return res
}