func findLUSlength(a string, b string) int {
    // uncommon [SUBSEQUENCE]
    
    if a == b {
        return -1
    }
    return max(len(a), len(b))
}