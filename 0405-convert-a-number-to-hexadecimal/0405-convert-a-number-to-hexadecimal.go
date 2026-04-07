func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    
    hexChars := "0123456789abcdef"
    var result []byte
    
    n := uint32(num)
    
    for n > 0 {
        digit := n & 15 
        result = append(result, hexChars[digit])
        n >>= 4       
    }
    
    for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
        result[i], result[j] = result[j], result[i]
    }
    
    return string(result)
}