/*
- Problem: https://leetcode.com/problems/longest-palindrome/
*/

func longestPalindrome(s string) int {
    
    var countMap map[rune]int = map[rune]int{}
    var doubleCharCount int
    
    for _, char := range s {
        if _, exists := countMap[char]; exists {
            countMap[char]++
          
            // Counting total double characters
            if countMap[char] % 2 == 0 {
                doubleCharCount++
            }
        } else {
            countMap[char] = 1
        }
    }
    
    if (doubleCharCount * 2) < len(s) {
        return (doubleCharCount * 2) + 1
    }
    
    return doubleCharCount * 2
    
}
