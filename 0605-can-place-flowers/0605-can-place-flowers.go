func canPlaceFlowers(flowerbed []int, n int) bool {
    cnt := 0
    length := len(flowerbed)

    // 길이가 1인 경우  
    if length == 1 {
        if flowerbed[0] == 0 {
            cnt++
        }
        return cnt >= n
    }

    // 맨 앞 처리
    if flowerbed[0] == 0 && flowerbed[1] == 0 {
        cnt++
        flowerbed[0] = 1
    }

    // 맨 뒤 처리
    if flowerbed[length-1] == 0 && flowerbed[length-2] == 0 {
        cnt++
        flowerbed[length-1] = 1
    }

    // 중간 sliding window
    for i := 1; i < length-1; i++ {
        // 가능 케이스
        if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
            cnt++
            flowerbed[i] = 1 
            i++              // 한 칸 건너뛰기
            continue    
        }

        // 왼쪽이 1
        if flowerbed[i-1] == 1 {
            continue
        }

        // 오른쪽이 1
        if flowerbed[i+1] == 1 {
            i++ // 한 칸 건너뛰기
            continue
        }
    }
    
    return cnt >= n
}