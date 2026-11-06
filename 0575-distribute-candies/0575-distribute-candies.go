func distributeCandies(candyType []int) int {
    n := len(candyType)

    uniqueCandies := make(map[int]bool)
    for _, v := range candyType {
        if !uniqueCandies[v] {
            uniqueCandies[v] = true
        }
    }

    if len(uniqueCandies) < n/2 {
        return len(uniqueCandies)
    } else {
        return n/2
    }
}