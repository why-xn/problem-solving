/*
- Problem: https://leetcode.com/problems/majority-element/
*/

class Solution {
    public int majorityElement(int[] nums) {
        
        HashMap<Integer, Integer> countMap = new HashMap<Integer, Integer>();
        int maxCount = 1;
        int maxN = nums[0];

        for (int n : nums) {
            if (countMap.containsKey(n)) {
                countMap.put(n, countMap.get(n)+1);
            } else {
                countMap.put(n, 1);
            }
            if (maxCount < countMap.get(n)) {
                maxCount = countMap.get(n);
                maxN = n;
            }
        }

        return maxN;
        
    }
}
