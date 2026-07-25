func maxVowels(s string, k int) int {
    vMap := map[byte]bool{
        'a' : true, 
        'e' : true, 
        'i' : true, 
        'o' : true, 
        'u' : true,
    }

    isVowel := func(s string, i int) int{
        if vMap[s[i]] == true {
            return 1
        }
        return 0
    }


    currentSum := 0
    // init Window
    for i := 0; i < k; i++ {
        if isVowel(s, i) == 1 {
            currentSum++
        }
    }

    maxSum := currentSum
    // mve sliding window
    for i := k; i < len(s); i++ {
        newVal := isVowel(s, i)
        prevVal := isVowel(s, i-k)

        currentSum = currentSum + newVal - prevVal

        if currentSum > maxSum {
            maxSum = currentSum
        }
    }

    return maxSum
}