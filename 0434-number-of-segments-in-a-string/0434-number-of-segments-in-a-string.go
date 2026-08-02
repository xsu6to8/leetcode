func countSegments(s string) int {
    res := strings.Fields(s)

    return len(res)
}