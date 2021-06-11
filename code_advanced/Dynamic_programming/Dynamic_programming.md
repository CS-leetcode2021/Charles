# Dynamic programming
-  动态规划的解题要点：   确定全局最优解、与最优子结构之间的关系、问题的边界 

---
##  背包问题

【思路连接】(https://leetcode-cn.com/leetbook/read/journey-of-algorithm/5200j3/)

---
##  金矿问题挖不挖的事情

-   每一个金矿都存在着「挖」和「不挖」两种选择，简化为两个最优子结构的问题(状态转移方程式)

    问题边界：金矿为0,工人为0的情况，那局部最优解就是0。
    **F(n,w)=0 (n≥1,w<p[n−1])**

    当所剩的工人不够挖当前金矿的时候只有一种最优子结构。
    **F(n,w)=F(n-1,w) (n≥1,w<p[n−1])**
    
    常规情况下，具有两种最优子结构(挖当前矿和不挖当矿)
    **F(n,w)=max(F(n−1,w),F(n−1,w−p[n−1])+g[n−1]) (n≥1,w≥p[n−1])**

    F(n−1,w) 不挖当前矿，F(n−1,w−p[n−1])+g[n−1] 挖当前矿