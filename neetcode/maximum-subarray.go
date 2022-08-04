/*
- Problem: https://leetcode.com/problems/maximum-subarray/
*/

func maxSubArray(nums []int) int {
    var r int = 0
	var currentSum int = 0
	var maxSum int = -2147483648

	for r < len(nums) {
		if currentSum+nums[r] < 0 {
			currentSum = 0
            if nums[r] > maxSum {
                maxSum = nums[r]
            }
		} else {
			currentSum = currentSum + nums[r]
			if currentSum > maxSum {
				maxSum = currentSum
			}
		}
        r++
	}

	return maxSum
}
