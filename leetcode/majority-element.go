/*
- Problem: https://leetcode.com/problems/majority-element/
*/

func majorityElement(nums []int) int {
    var trackMap map[int]int = map[int]int{}
    var nlen int = len(nums)
    var maxV int = 0
    var maxK int = 0
    
    for _, n := range nums {
        if _, ok := trackMap[n]; ok {
            trackMap[n] = trackMap[n] + 1
        } else {
            trackMap[n] = 1
        }
    }
    
    for k, v := range trackMap {
        if v >= nlen / 2 {
            if maxV < v {
                maxV = v
                maxK = k
            }
        }
    }
    
    return maxK
}
