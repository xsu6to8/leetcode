/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    start := 1
    end := n
    
    for start < end {
        // overflow prevetntion
        mid := start + (end - start) / 2
        
        if isBadVersion(mid) {
            // 현재 mid가 Bad, 정답은 mid이거나 mid 왼쪽에 있음
            end = mid
        } else {
            // 현재 mid가 good, 정답은 무조건 mid+1부터 있음
            start = mid + 1
        }
    }
    
    return start
}
