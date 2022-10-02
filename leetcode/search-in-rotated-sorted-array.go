/*
- Problem: https://leetcode.com/problems/search-in-rotated-sorted-array/
*/

func search(nums []int, target int) int {
    f := nums[0]
    ri := 0
    
    // Fixing the rotation and making it sorted again
    if nums[0] > nums[len(nums)-1] {
        for i:=len(nums) - 1; i > 0; i-- {
            if nums[i] > nums[0] {
                break
            }
            ri++
        }
        
        tnums := []int{}
        tnums = append(tnums, nums[len(nums) - ri:len(nums)]...)
        tnums = append(tnums, nums[0:len(nums) - ri]...)
        
        nums = tnums
        
    } else {
        ri = 0
    }
    
    res := binarySearch(len(nums)/2,  0, len(nums) - 1, nums, target)
    
    // The res returns index based on the sorted array. Re-calculating the index based on rotated array.
    if res >= 0 {
        if nums[res] < f {
            res = res + (len(nums) - ri)
        } else if nums[res] >= f {
            res = res - ri
        }
    }
    
    return res
}

func binarySearch(curr int, min int, max int, nums []int, target int) int {
    if nums[curr] == target {
        return curr
    } else if target < nums[curr] {
        if curr - min == 0 {
            return -1
        } else {
            return binarySearch(((curr + min)/2), min, curr - 1, nums, target)
        }
    } else if target > nums[curr] {
        if max - curr == 0 {
            return -1
        } else {
            return binarySearch(((curr + max)/2) + 1, curr + 1, max, nums, target)
        }
    }
    return -1
}
