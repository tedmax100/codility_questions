[123. Best Time to Buy and Sell Stock III](https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/)

<code>
You are given an array prices where prices[i] is the price of a given stock on the ith day.  

Find the maximum profit you can achieve. You may complete at most two transactions.  

Note: You may not engage in multiple transactions simultaneously (i.e., you must sell the stock before you buy again).  

 

Example 1:  
Input: prices = [3,3,5,0,0,3,1,4]  
Output: 6  
Explanation: Buy on day 4 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.
Then buy on day 7 (price = 1) and sell on day 8 (price = 4), profit = 4-1 = 3.  

Example 2:  
Input: prices = [1,2,3,4,5]  
Output: 4  
Explanation: Buy on day 1 (price = 1) and sell on day 5 (price = 5), profit = 5-1 = 4.
Note that you cannot buy on day 1, buy on day 2 and sell them later, as you are engaging multiple transactions at the same time. You must sell before buying again.  

Example 3:  
Input: prices = [7,6,4,3,1]  
Output: 0  
Explanation: In this case, no transaction is done, i.e. max profit = 0.  

Constraints:  
1 <= prices.length <= 10<sup>5</sup>  
0 <= prices[i] <= 10<sup>5</sup>
</code>

<details>
  <summary>中文意思</summary>

> 給一組prices int array, 其中prices[i]表示當天股票的價格.  
> 目標將利潤給最大化, 但規定過程中最多執行2次的交易, 且手上最多也只能有一筆股票 
> 如果無法獲得利潤, 則返回0  

# 買低賣高!  
粗淺想法是找出最高效益的跟次高的作加總  
延續前面的作法, 今天比昨天高就賣出的策略, 發現有反例  
```[1,5,2,8,3,10]```, 三次買賣分別利潤是```4,6,7```, 取前兩高加總 = 13  
但其實Day1能等到Day4的8再賣出, 利潤分別是```7,7```, 加總是14  
所以買賣策略沒那麼單純了!  

所以換個思路, 按照DP的思維, 拆成兩堆profit[i]表示前i天的最高立論跟profit[i+1]  
加上最多交易2次的這條件, 交易次數用```k```表示  
profit[i][k]表示前i天最多k次的最高收益  

每一天都能選擇買賣或不做事, 如果在第i天選擇賣出, 那也就表示, 在0 ~ i-1添有一天要選擇買入, 才能做這賣出操作, 也表示可能在這買入之前也進行了k-1次買賣.  
第1天就買入, 第i天做賣出, 利潤是 prices[i] - prices[0]  
第2天買入, 第i天做賣出, 利潤是 prices[i] - prices[1] + profit[0][k-1] , 加上前1天的最大收益  
第3天買入, 第i天做賣出, 利潤是 prices[i] - prices[2] + profit[1][k-1] , 依此類推...  
第j天買入, 第i天做賣出, 利潤是 prices[i] - prices[j] + profit[j-1][k-1]  
在這些可能裡面選擇一個最大的, 跟第i天什麼都不做事做比較, 就是profit[i][k]的值了  

https://zhuanlan.zhihu.com/p/77666061
https://leetcode.com/problems/best-time-to-buy-and-sell-stock-iii/discuss/149383/Easy-DP-solution-using-state-machine-O(n)-time-complexity-O(1)-space-complexity
</details>