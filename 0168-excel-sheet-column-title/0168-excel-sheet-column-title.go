func convertToTitle(columnNumber int) string {
    res := ""
    
    for columnNumber > 0 {
        // 0-indexed 변환 (A=0, Z=25)
        columnNumber-- 
        
        char := string(rune('A' + (columnNumber % 26)))
        
        res = char + res
        
        columnNumber /= 26
    }
    
    return res
}