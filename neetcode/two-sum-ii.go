/*
- Problem: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/
*/

func twoSum(numbers []int, target int) []int {
    
    var i int = 0
    var j int = len(numbers) - 1
    var diff int = 0
    
    for i < j {
        diff = target - (numbers[i] + numbers[j])
        if diff == 0 {
            return []int{i+1, j+1}
        } else if diff > 0 {
            i++
        } else if diff < 0 {
            j--
        }
    }
    return []int{}
}
