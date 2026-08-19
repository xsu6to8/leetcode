func findRelativeRanks(score []int) []string {
    // deep copy : ONLY values
    origin := make([]int, len(score))
    copy(origin, score)

    sort.Slice(score, func(i, j int) bool {
		return score[i] > score[j]
	})

    sMap := make(map[int]int)
    for i, v := range score {
        sMap[v] = i+1
    }

    var res []string
    for _, v := range origin {
        rank := sMap[v]
        if rank == 1 {
            res = append(res, "Gold Medal")
        } else if rank == 2 {
            res = append(res, "Silver Medal")
        } else if rank == 3 {
            res = append(res, "Bronze Medal")
        } else {
            res = append(res, strconv.Itoa(rank))
        }
    }

    return res
}