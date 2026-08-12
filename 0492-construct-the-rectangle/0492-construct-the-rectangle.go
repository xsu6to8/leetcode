// func constructRectangle(area int) []int {
//     sqrt := int(math.Sqrt(float64(area)))

//     var res []int
//     gap := math.MaxInt
//     for i := 1; i <= sqrt; i++ {
//         if area % i == 0 {
//             v1 := area / i
//             v2 := i

//             maxV := max(v1, v2)
//             minV := min(v1, v2)
//             tmpGap := v1 - v2

//             if tmpGap < gap {
//                 res = res[:0]
//                 res = append(res, maxV, minV)
//             }
//         }
//     }

//     return res
// }


// Optimal solution
func constructRectangle(area int) []int {
	// sqrt부터 역순으로 탐색 -> gap이 가장 작음이 확실시 [나누는 값이 크기에]
	for w := int(math.Sqrt(float64(area))); w >= 1; w-- {
		if area%w == 0 {
			// 조건 충족하는 최초로 발견한 (area/w, w)가 차이가 가장 작은 조합
			return []int{area / w, w}
		}
	}
	return nil
}