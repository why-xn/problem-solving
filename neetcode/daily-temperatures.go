/*
- Problem: https://leetcode.com/problems/daily-temperatures/
*/

func dailyTemperatures(temperatures []int) []int {
    var res []int = make([]int, len(temperatures))
    var iStack []int = []int{} // A stack for storing the indices of the temperatures for which no warmer day has been found yet.
    var slen int = 0
    
    // appending the first temperature in the stack by default
    iStack = append(iStack, 0)
    slen++
    
    for i := 1; i < len(temperatures); i++ {
        if slen > 0 && temperatures[i] > temperatures[iStack[slen - 1]] {
            // if a warmer temperature is found than the top of the stack then pop all the cooler temperatures and calculate their results
            for slen > 0 {
                if temperatures[i] <= temperatures[iStack[slen - 1]] {
                    break
                } else {
                    res[iStack[slen-1]] = i - iStack[slen-1]
                    iStack = iStack[:slen-1]
                    slen--
                }
            }
            
        }
        iStack = append(iStack, i)
        slen++
    }
    
    return res
    
}
