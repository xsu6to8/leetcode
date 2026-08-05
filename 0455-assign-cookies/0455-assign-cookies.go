func findContentChildren(g []int, s []int) int {
    sort.Ints(g)
    sort.Ints(s)

    gIdx := 0
    sIdx := 0

    for gIdx < len(g) && sIdx < len(s) {
		if s[sIdx] >= g[gIdx] {
			gIdx++ 
		}
		sIdx++
	}

	return gIdx
}