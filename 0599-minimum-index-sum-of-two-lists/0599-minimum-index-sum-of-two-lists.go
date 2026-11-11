func findRestaurant(list1 []string, list2 []string) []string {
    if len(list1) > len(list2) {
        list1, list2 = list2, list1
    }
    
    listMap := make(map[string]int)
    for i, v := range list1 {
        listMap[v] = i
    }

    var res []string
    sumVal := len(list1) + len(list2)
    for i, v := range list2 {
        if j, ok := listMap[v]; !ok {
            continue
        } else {
            if i + j > sumVal {
                continue
            } else if i + j == sumVal {
                res = append(res, v)
            } else {
                sumVal = i + j
                res = res[:0]
                res = append(res, v)
            }
        }

    }

    return res
}