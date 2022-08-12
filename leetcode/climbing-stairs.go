/*
- Problem: https://leetcode.com/problems/climbing-stairs/
*/

func climbStairs(n int) int {
    
    if n < 3 {
        return n
    }
    
    cache := make([]int, n+1)
    
    cache[n] = 1
    cache[n-1] = 1
    
    // Bottom up approach to calculate possible ways from each position. Possible Ways From N = (Possible Way From N + 1) + (Possible Way From N + 2)
    for j:=n-2; j >=0;  j-- {
        cache[j] = cache[j+1] + cache[j+2]
    }
    
    return cache[0]
    
}

