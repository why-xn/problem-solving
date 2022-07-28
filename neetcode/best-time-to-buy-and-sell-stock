*/
- https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
*/

func maxProfit(prices []int) int {
    
    var tmpMin int = prices[0]
    var maxProfit int = 0
    
    for _, p := range prices {
        if p - tmpMin > maxProfit {
            maxProfit = p - tmpMin
        }
        if p < tmpMin {
            tmpMin = p
        }
    }
    
    return maxProfit
    
}
