/*
	给定不同面额的硬币 coins 和一个总金额 amount。

	编写一个函数来计算可以凑成总金额所需的最少的硬币个数。

	如果没有任何一种硬币组合能组成总金额，返回 -1。
*/
func coinChange(coins []int, amount int) int {
	//dp[i]表示总金额为i时的硬币个数
	//dp[i]=min(dp[i-coins[_]]+1,dp[i]]
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
		for _, coin := range coins {
			if i < coin || dp[i-coin] == -1 {
				continue
			}
			count := dp[i-coin] + 1
			if dp[i] == -1 || dp[i] > count {
				dp[i] = count
			}
		}
	}
	return dp[amount]
}