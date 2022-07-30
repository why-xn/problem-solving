/*
- Problem: https://leetcode.com/problems/permutation-in-string/
*/

func checkInclusion(s1 string, s2 string) bool {
    
    var srcHash map[byte]int = map[byte]int{}
    var matchHash map[byte]int = map[byte]int{} // to store frequecy of characters of currency sliding window
    var s1len int = len(s1)
    
    // saving character frequency of s1 in a hash map
    for i:=0; i<s1len; i++ {
        if _, ok := srcHash[s1[i]]; ok {
            srcHash[s1[i]] = srcHash[s1[i]] + 1
        } else {
            srcHash[s1[i]] = 1
        }
    }
    
    var l, r , length int = 0, 0, 0
    
    for r < len(s2) {
        if _, ok := srcHash[s2[r]]; ok {
            if matchHash[s2[r]] == srcHash[s2[r]] {
                // if character from s2 exists in s1 and the frequency of them also matches already, then adding this to sliding window will create a mismatch.
                // so removing all the characters from sliding window until the first occurance of the matching character is removed and creating a room for adding the current one.
                for {
                    if s2[l] != s2[r] {
                        matchHash[s2[l]] = matchHash[s2[l]] - 1
                        l++
                        length--
                    } else {
                        l++
                        break
                    } 
                }
            } else {
                // increasing frequency count for a match
                matchHash[s2[r]] = matchHash[s2[r]] + 1
                length++
                if length == s1len {
                    return true
                }
            }
        } else {
            // if character does not exists in s1 then reset the sliding window
            length = 0
            matchHash = map[byte]int{}
            l = r + 1
        }
        r++
    }
    
    if length == s1len {
        return true
    }
    
    return false
}
