func longestPalindrome(s string) int {
    // alphbet 65~ / 97~ : 128 크기의 배열
    var counts [128]int
    
    for _, char := range s {
        counts[char]++
    }
    
    length := 0
    hasOdd := false
    
    for _, count := range counts {
        // cnt 보다 작거나 같은 짝수 전부
        length += (count / 2) * 2
        
        // 홀수인 문자가 하나라도 있는지 체크
        if count%2 == 1 {
            hasOdd = true
        }
    }
    
    // 홀수 문자가 하나라도 있었다면 길이에 1 추가
    if hasOdd {
        length++
    }
    
    return length
}