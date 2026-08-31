func search(nums []int, target int) int {
    if len(nums) == 1 {
        if nums[0] == target {
            return 0
        } else {
            return -1
        }
    }
    
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l + r) / 2

        // binary search로 찾음
        if nums[mid] == target {
            return mid
        }

        if nums[l] <= nums[mid] {
        // case1 : ORDERED left parts
            // target이 정렬된 왼쪽 구간에 존재 가능 -> 왼쪽으로 범위 축소
            if nums[l] <= target && target <= nums[mid] {
                r = mid - 1
            } else { // 없다면 -> 오른쪽 구간 탐색
                l = mid + 1
            }
        } else {
        // case2 : ORDERED right parts
            // target이 정렬된 오른쪽 구간에 존재 가능 -> 오른쪽으로 범위 축소
            if nums[mid] <= target && target <= nums[r] {
                l = mid + 1
            } else { // 없다면 -> 왼쪽 구간 탐색
                r = mid - 1
            }
        }
    }

    return -1
}