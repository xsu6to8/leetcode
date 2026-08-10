func findComplement(num int) int {
    if num == 1 {
        return 0
    } 
    
    nearestPowOfTwo := 1
    for num >= nearestPowOfTwo {
        nearestPowOfTwo *= 2
    }
    nearestPowOfTwo -= 1

    return nearestPowOfTwo ^ num
}