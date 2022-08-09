/*
- Problem: https://leetcode.com/problems/is-subsequence/
*/

func isSubsequence(s string, t string) bool {
    
    var i, j int = 0, 0
    
    if len(s) > 0 {
        for j < len(t) {
            if s[i] == t[j] {
                i++
                if i == len(s) {
                    return true
                }
            }
            j++
        }
    }
    
    if i == len(s) {
        return true
    }
    
    return false
    
}
