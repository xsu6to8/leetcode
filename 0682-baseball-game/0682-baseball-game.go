func calPoints(operations []string) int {
    var st []int
    for _, v := range operations {
        if v == "C" ||  v == "D" ||  v == "+" {
            if v == "C" {
                st = st[:len(st)-1]
            } else if v == "D" {
                curr := st[len(st)-1]
                st = append(st, curr*2)
            } else if v == "+" {
                fir, sec := st[len(st)-1], st[len(st)-2]
                st = append(st, fir+sec)
            }
        } else {
            vInt, _ := strconv.Atoi(v)
            st = append(st, vInt)
        }
    }

    total := 0
	for _, n := range st {
		total += n
	}
	return total
}