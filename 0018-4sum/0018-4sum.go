func fourSum(nums []int, target int) [][]int {
    slices.Sort(nums)
    n := len(nums)

    res := [][]int{}

    for i := 0; i < n-3; i++ {
        // 1. 현재 가능한 가장 작은 4개의 합이 target보다 크면, 이후는 볼 필요 없음
        if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target { break } 
        // 2. 현재 i와 가장 큰 3개의 합이 target보다 작으면, 현재 i는 건너뜀
        if nums[i]+nums[n-3]+nums[n-2]+nums[n-1] < target { continue }

        // 초깃값 이전 것과 동일하면 넘어가기
        if i > 0 && nums[i] == nums[i-1] { continue }

        for j := i+1; j < n-2; j++ {
            // 초깃값의 [다음 값] 이전 것과 동일하면 넘어가기
            if j > i+1 && nums[j] == nums[j-1] { continue }

            left := j + 1
            right := n - 1

            for left < right {
                sum := nums[i] + nums[j] + nums[left] + nums[right]

                if sum == target {
                // 정답 추가
                res = append(res, []int{nums[i], nums[j], nums[left], nums[right]})
        
                // left 포인터 이동 (중복 제거 포함)
                for left < right && nums[left] == nums[left+1] {
                    left++
                }
                // right 포인터 이동 (중복 제거 포함)
                for left < right && nums[right] == nums[right-1] {
                    right--
                }
        
                // 정답을 찾았으니 한 칸씩 더 이동
                left++
                right--
                } else if sum < target {
                    left++
                } else {
                    right--
                }
            }  
        }
    }
    return res
}