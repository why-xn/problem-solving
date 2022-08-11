/*
- Problem: https://leetcode.com/problems/3sum/
*/

func threeSum(nums []int) [][]int {
    
    var hset map[string]*byte = map[string]*byte{}
    var res [][]int = [][]int{}
    var i int = 0
    var j int = len(nums) - 1
    var diff int = 0
    var hash string = ""
    
    // Sorting the array
    sort.Ints(nums)
    
    
    for x, n := range nums {
        i = x + 1;
        j = len(nums) - 1
        
        // Applying 2sum
        for i < j {
            diff = 0 - (n + nums[i] + nums[j])
            if diff == 0 {
                hash = fmt.Sprintf("%d-%d-%d", n, nums[i], nums[j])
                if _, ok := hset[hash]; ok {
                    i++
                    j--
                    continue
                }
                hset[hash] = nil
                res = append(res, []int{n, nums[i], nums[j]})
                i++
                j--
            } else if diff > 0 {
                i++
            } else if diff < 0 {
                j--
            }
        }
    }
    
    return res
    
}
