/*
- Problem: https://leetcode.com/problems/first-bad-version/
*/

/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
    
    var first int = 0
    var last int = n
    var mid int
    var fbv int
    
    // doing binary search
    for first <= last {
        mid = (first + last) / 2
        
        if isBadVersion(mid) {
            last = mid -1
            fbv = mid
        } else {
            first = mid + 1
        }
    }
    
    return fbv
    
}
