/*
- Problem: https://leetcode.com/problems/longest-consecutive-sequence/
*/

func longestConsecutive(nums []int) int {
    var nset map[int]*byte = map[int]*byte{}
    
    // adding all numbers to set
    for _, n := range nums {
        nset[n] = nil
    }
    
    var maxSequence int = 0
    var length int = 0
    
    for _, n := range nums {
        
        // if previous number does not exists, its a new sequence
        if _, ok := nset[n - 1]; !ok {
            length = 0
            for true {
                if _, ok := nset[n + length]; ok {
                    length++
                } else {
                    break
                }
            }
            if maxSequence < length {
                maxSequence = length
            }
        }
    }
    return maxSequence   
}
