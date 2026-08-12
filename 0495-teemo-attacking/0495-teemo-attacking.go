func findPoisonedDuration(timeSeries []int, duration int) int {
    total := 0
    for i := 0; i < len(timeSeries); i++ {
        currT := timeSeries[i]
        endT := timeSeries[i] + duration - 1
        for i + 1 < len(timeSeries) && timeSeries[i + 1] <= endT {
            endT = timeSeries[i + 1] + duration - 1
            i++
        }
        total += endT - currT + 1
    }

    return total
}