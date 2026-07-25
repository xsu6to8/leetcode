func uniqueOccurrences(arr []int) bool {
    mySet := make(map[int]struct{})
    
    sort.Ints(arr)
    
    currVal := arr[0]
    cnt := 0
    for i := 0; i < len(arr); i++ {
        if currVal != arr[i] {
            if _, exists := mySet[cnt]; exists {
                return false
            }
            
            mySet[cnt] = struct{}{}
            
            currVal = arr[i]
            cnt = 1
            continue
        }
        cnt++
    }

    // 마지막 idx
    if _, exists := mySet[cnt]; exists {
                return false
    }

    return true
}