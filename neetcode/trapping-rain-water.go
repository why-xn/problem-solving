/*
- Problem: https://leetcode.com/problems/trapping-rain-water/
*/

func trap(height []int) int {
    var maxL int = 0
    var maxR int = height[len(height)-1]
    
    var res int = 0
    
    var i = 0
    var j = len(height) - 1
    var c int = i
    var trap int = 0
    for i < j {
        trap = getMin(maxL, maxR) - height[c]
        if (trap > 0) {
            res = res + trap
        }
        if c == i && height[c] > maxL {
            maxL = height[c]
        } else if c == j && height[c] > maxR {
            maxR = height[c]
        }
        if height[i] < height[j] {
            i++
            c = i
        } else {
            j--
            c = j
        }
    }
    return res
}

func getMin(f int, s int) int {
    if f < s {
        return f
    } else {
        return s
    } 
}
