func maxArea(height []int) int {
    left := 0
    right := len(height) - 1
    res := -1

    for left < right{
        currS := (min(height[left], height[right])) * (right - left)
        if currS > res {
            res = currS
        }       

        if height[left] == min(height[left], height[right]) {
            left++
        } else {
            right--
        }
    }

    return res
}

func min(a, b int) int {
    if a < b {
        return a
    } else {
        return b
    }
}