/*
- Problem: https://leetcode.com/problems/find-pivot-index/
*/

func pivotIndex(nums []int) int {
    
    var numLength int = len(nums)
    var prefixSum []int = make([]int, numLength)
    var reversePrefixSum []int = make([]int, numLength)
    
    var lSum, rSum int = 0, 0
    
    for i:=0; i<numLength;i++ {
        prefixSum[i] = lSum
        lSum = lSum + nums[i]
        
        reversePrefixSum[numLength - 1 - i] = rSum
        rSum = rSum + nums[numLength - 1 - i]
    }
    
    for i:=0; i<numLength;i++ {
        if prefixSum[i] == reversePrefixSum[i] {
            return i
        }
    }
    return -1
}
