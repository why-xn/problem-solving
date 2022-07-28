func isPalindrome(s string) bool {
    var i int = 0
    var j int = len(s) - 1
    
    s = strings.ToLower(s)
    
    for i < j {
        if ((s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57)) && ((s[j] >= 97 && s[j] <= 122) || (s[j] >= 48 && s[j] <= 57)) {
            if s[i] != s[j] {
                return false
            } else {
                i++
                j--
            }
        } else {
            if !((s[i] >= 97 && s[i] <= 122) || (s[i] >= 48 && s[i] <= 57)) {
                i++
            }
            if !((s[j] >= 97 && s[j] <= 122) || (s[j] >= 48 && s[j] <= 57)) {
                j--
            }
        }
    }
    return true
}
