// func findLHS(nums []int) int {
//     maxL := 0
    
//     slices.Sort(nums)
//     l, r := 0, 0

//     for l < len(nums) {
//         cnt := 0
//         curr := nums[l]
//         for l + 1 < len(nums) && curr == nums[l+1] {
//            cnt++
//            l++ 
//         }
//         r = l + 1
//         if r < len(nums) {
//             next := nums[r]
//             if next == curr + 1 {
//                 for r + 1 < len(nums) && next == nums[r+1] {
//                     cnt++
//                     r++
//                 }
//                 if cnt +2 > maxL {
//                     maxL = cnt+2
//                 }
//             }
//         } 
//         l++
//     }

//     return maxL
// }

// method2 : Map 사용
func findLHS(nums []int) int {
    nMap := make(map[int]int)
    for _, v := range nums {
        if _, ok := nMap[v]; !ok {
            nMap[v] = 1
        } else {
            nMap[v]++
        }
    }

    maxL := 0
    for i, v := range nMap {
        if _, ok := nMap[i+1]; !ok {
            continue
        } else {
            currL := v + nMap[i+1]
            if currL > maxL {
                maxL = currL
            }
        }
    }

    return maxL
}