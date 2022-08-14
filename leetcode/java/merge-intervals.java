/*
- Problem: https://leetcode.com/problems/merge-intervals/
*/

class Solution {
    public int[][] merge(int[][] intervals) {
        Arrays.sort(intervals, (a, b) -> Integer.compare(a[0], b[0]));
        List<int[]> result = new ArrayList<int[]>();

        int[] mergingArray = intervals[0];

        for (int i = 1; i < intervals.length; i++) {
            if (mergingArray[1] < intervals[i][0]) {
                // Breaking the current merge and starting a new one.
                result.add(mergingArray);
                mergingArray = intervals[i];
            } else {
                // Increasing the merging interval length if required.
                mergingArray[1] = intervals[i][1] > mergingArray[1] ? intervals[i][1] : mergingArray[1];
                continue;
            }
        }

        result.add(mergingArray);

        return result.stream().toArray(int[][]::new);
    }
}
