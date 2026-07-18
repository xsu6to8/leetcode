func kidsWithCandies(candies []int, extraCandies int) []bool {
    maxCandies := slices.Max(candies)

    result :=  make([]bool, len(candies))

    for i, v := range(candies) {
        if v + extraCandies >= maxCandies {
            result[i] = true
        } else {
            result[i] = false
        }
    }

    return result
}