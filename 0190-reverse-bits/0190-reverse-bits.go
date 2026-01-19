func reverseBits(n int) int {
    var result int = 0
    
    for i := 0; i < 32; i++ {
        result = result << 1
        
        bit := n & 1
        result = result | bit

        n = n >> 1
    }
    
    return result
}