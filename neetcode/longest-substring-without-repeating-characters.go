/*
- Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
    
    var hash map[byte]*byte = map[byte]*byte{}
    
    var l int = 0
    var r int = 0
    var max int = 0
    
    for r < len(s) {
        if _, ok := hash[s[r]]; ok {
            if max < len(hash) {
                max = len(hash)
            }
            for {
                delete(hash, s[l])
                if s[l] == s[r] {
                    l++
                    break;
                }
                l++
            }
            hash[s[r]] = nil
        }
        
        hash[s[r]] = nil
        r++
    }
    
    if max < len(hash) {
        max = len(hash)
    }
    
    return max
    
}
