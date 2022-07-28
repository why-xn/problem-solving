/*
- Problem: https://leetcode.com/problems/longest-substring-without-repeating-characters/
*/

func lengthOfLongestSubstring(s string) int {
    
    var hash map[byte]*byte = map[byte]*byte{}
    
    var l int = 0
    var r int = 0
    var max int = 0
    
    for r < len(s) {
        // if character exists in hash, 
        if _, ok := hash[s[r]]; ok {
            // as adding new character will create duplicacy in the hash, this is the max sequence
            if max < len(hash) {
                max = len(hash)
            }
            
            // pop all the characters from hash until the matched character is removed from hash to avoid duplicacy
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
        
        // add character to hash
        hash[s[r]] = nil
        r++
    }
    
    // final check
    if max < len(hash) {
        max = len(hash)
    }
    
    return max
    
}
