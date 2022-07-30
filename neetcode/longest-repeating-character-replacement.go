/*
- Problem: https://leetcode.com/problems/longest-repeating-character-replacement/
*/

func characterReplacement(s string, k int) int {
    
    var current byte = s[0]
    var i int = 1
    var checkpoint int = 0
    var max int = 1
    var remainingReplaceCount int = k
    var length int = 1
    
    for i < len(s) {
        // if current character is equal to previous character, then increasing the sequence length
      
        if s[i] != current && remainingReplaceCount == 0 {
            // if doesn't match and no replace option available then start a new sequence
            if length > max {
                max = length
            }
            if remainingReplaceCount < k {
                i = checkpoint
            }
            current = s[i]
            length = 0
            remainingReplaceCount = k
            
        } else if s[i] != current && remainingReplaceCount > 0 {
            // if doesn't match and replace option available then store the first replacement index as checkpoint
            if remainingReplaceCount == k {
                checkpoint = i
            }
            remainingReplaceCount--
        }
        i++
        length++
    }
    
    // if replacement count still available after iterating through all the last sequence, remaining replacements are availed from the left side of the last sequence. 
    if remainingReplaceCount > 0 {
        if remainingReplaceCount <= len(s) - length {
            length = length + remainingReplaceCount
        } else {
            length = length + ((len(s)) - length)
        }
    }
    
    // checking if the last sequence is max or not
    if length > max {
        max = length
    }
    
    return max
    
}
