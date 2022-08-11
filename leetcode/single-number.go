/*
- Problem: https://leetcode.com/problems/single-number/
*/

func singleNumber(nums []int) int {
    
    var trackMap map[int]int = map[int]int{}
    
    for _, n := range nums {
        if _, ok := trackMap[n]; ok {
            trackMap[n] = trackMap[n] + 1
        } else {
            trackMap[n] = 1
        }
    }
    
    for k, v := range trackMap {
        if v == 1 {
            return k
        }
    }
    
    return 0
    
}
