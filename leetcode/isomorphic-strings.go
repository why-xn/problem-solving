/*
- Problem: https://leetcode.com/problems/isomorphic-strings/
*/


func isIsomorphic(s string, t string) bool {
    
    var sMap, tMap map[byte]int = map[byte]int{}, map[byte]int{}
    var sHash, tHash int = 0, 0
    var sIndx, tIndx int = 0, 0
    
    for i:=0; i<len(s); i++ {
        // Setting a unique number for each character in the (s) string based on the order of it's first occurence and calculate the hash of the string
        if val, ok := sMap[s[i]]; ok {
            sHash = (sHash * 10) + val
        } else {
            sIndx++
            sMap[s[i]] = sIndx
            sHash = (sHash * 10) + sIndx
        }
        
        // Setting a unique number for each character in the (t) string based on the order of it's first occurence and calculate the hash of the string
        if val, ok := tMap[t[i]]; ok {
            tHash = (tHash * 10) + val
        } else {
            tIndx++
            tMap[t[i]] = tIndx
            tHash = (tHash * 10) + tIndx
        }
    }
    
    // Checking if hashes of both string matches or not
    if sHash == tHash {
        return true
    }
    
    return false
    
}
