/*
- Problem: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
*/

func searchRange(nums []int, target int) []int {
    
    if len(nums) == 0 {
        return []int{-1,-1}
    }
    
    mid := len(nums) / 2
    
    // finding if the number exists or not, if exists get any of it's index
    idx := binarySearch(mid,  0, len(nums) - 1, nums, target)
    
    
    // find the first index
    first := idx
    for i := first - 1; i >= 0; i-- {
        if nums[i] == target {
            first = i
        } else {
            break;
        }
    }
    
    // find the last index
    last := idx
    for i := last + 1; i < len(nums); i++ {
        if nums[i] == target {
            last = i
        } else {
            break;
        }
    }
    
    return []int{first, last}
    
}

func binarySearch(curr int, min int, max int, nums []int, target int) int {
    if nums[curr] == target {
        return curr
    } else if target < nums[curr] {
        if curr - min == 0 {
            return -1
        } else if curr - min == 1 {
            return binarySearch(min, min, curr - 1, nums, target)
        } else {
            return binarySearch(((curr + min)/2), min, curr - 1, nums, target)
        }
    } else if target > nums[curr] {
        if max - curr == 0 {
            return -1
        } else if max - curr == 1 {
            return binarySearch(max, curr + 1, max, nums, target)
        } else {
            return binarySearch(((curr + max)/2), curr + 1, max, nums, target)
        }
    }
    return -1
}
