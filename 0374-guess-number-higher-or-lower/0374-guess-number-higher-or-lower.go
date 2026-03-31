/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */

func guessNumber(n int) int {
    l, r := 0, n
    mid := (l + r) / 2
    for {
        curr := guess(mid)
        switch curr {
            case -1 :
                r = mid - 1
                mid = (l + r) / 2
            case 1 :
                l = mid + 1
                mid = (l + r) / 2
            case 0 :
                return mid
        }
    }
}