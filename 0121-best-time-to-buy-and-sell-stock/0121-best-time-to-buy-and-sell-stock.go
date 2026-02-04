func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }

    minPrice := prices[0]
    maxProfit := 0
    
    for i := 1; i < len(prices); i++ {
        if prices[i] < minPrice {
            minPrice = prices[i]
        }

        if maxProfit < prices[i] - minPrice {
            maxProfit = prices[i] - minPrice
        }
    } 

    return maxProfit
}