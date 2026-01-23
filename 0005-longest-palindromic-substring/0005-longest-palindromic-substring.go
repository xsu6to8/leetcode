func longestPalindrome(s string) string {
    n := len(s)
    
    if n < 2 {
        return s
    }

    // dp[i][j] : s[i...j]가 팰린드롬인지
    dp := make([][]bool, n)
    for i := range dp {
        dp[i] = make([]bool, n)
    }

    start := 0
    maxLen := 1

    // #1. 길이가 1 == 단일 문자는 그 자체로 팰린드롬
    for i := 0; i < n; i++ {
        dp[i][i] = true
    }

    // length: 검사할 부분 문자열의 길이 (2부터 n까지)
    for length := 2; length <= n; length++ {
        // i: 시작 인덱스
        for i := 0; i <= n-length; i++ {
            j := i + length - 1 // j: 끝 인덱스

            // 양 끝 문자가 같고
            if s[i] == s[j] {
                // 길이가 2        -> 무조건 True
                // 길이가 3 이상    -> 안쪽(i+1...j-1)도 팰린드롬이어야 함
                if length == 2 || dp[i+1][j-1] {
                    dp[i][j] = true
                    
                    // 최대 길이 갱신
                    if length > maxLen {
                        start = i
                        maxLen = length
                    }
                }
            }
        }
    }

    return s[start : start+maxLen]
}