/*
- Problem: https://leetcode.com/problems/merge-intervals/
*/

func merge(intervals [][]int) [][]int {
    
    // First, sorting the 2d array based on the first index of each sub array
    sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })
    
    var result [][]int = [][]int{}
  
    // Initializing first merging interval
    var tmp []int = intervals[0]
    
    for _, intvl := range intervals {
        // If the last value of merging interval is more than the first value of the current interval, then these 2 needs to be merged. 
        // Or here breaks the current merging interval and starts a new one.
        if tmp[1] < intvl[0] {
            result = append(result, tmp)
            tmp = intvl
        } else {
            if tmp[1] < intvl[1] {
                tmp[1] = intvl[1]
                continue
            }
        }
    }
    
    result = append(result, tmp)
    return result   
}
