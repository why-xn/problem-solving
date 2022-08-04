/*
- Problem: https://leetcode.com/problems/fibonacci-number/
*/

func fib(n int) int {
    var fmap map[int]int = map[int]int{
        0: 0,
        1: 1,
    }
    
    var i int = 2
    for i <= n {
        fmap[i] = fmap[i-1] + fmap[i-2]
        i++
    }
    
    return fmap[n]
}
