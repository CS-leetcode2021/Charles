/*矩阵转置：按照对角线进行翻转*/

/*
	1	4	7
	2	5	8
	3	6	9
*/
------------------------------------------------------------------------------------------------------------------------
// #876 transpose 转置矩阵
![](https://assets.leetcode.com/uploads/2021/02/10/hint_transpose.png)
```go
	func transpose(matrix [][]int) [][]int {
		m, n := len(matrix), len(matrix[0])
		ns := make([][]int, m)
		for i := range ns {
			ns[i] = make([]int, n)
			for j := range ns[i] {
				ns[i][j] = -1
			}
		}
	
		for i,row := range matrix{
			for j , v := range row{
				ns[j][i]=v
			}
		}
		return ns
	}
```

```java
	class Solution {
	public int[][] transpose(int[][] matrix) {
		int m = matrix.length;
		int n = matrix[0].length;

		int[][] ns = new int[n][m];

		for (int i=0;i<m;i++){
			for (int j = 0; j<n;j++){
				ns[j][i] = matrix[i][j];
			}
		}
		return ns;
	}
}
```

------------------------------------------------------------------------------------------------------------------------

# 思路

-  采用摩尔投票法
-  对抗阶段：分属两个候选人的票数进行两两对抗抵消
-  计数阶段：计算对抗结果中最后留下的候选人票数是否有效
   [摩尔投票法讲解](https://cloud.tencent.com/developer/article/1600607)
- 最后要添加对众数是否过半的判定（!!!补充，摩尔投票法仅仅是对众数的判断，这里需要添加时候统计过半的条件）

```go
		func majorityElement(nums []int) int {
			target := 0
			count := 0
		
			for  i:= 0 ; i<len(nums);i++{
				// 如果当前计数为0,则使当前值为target
				if count == 0{
					target = nums[i]
				}
				// 不同则减一，相同则加一
				if target == nums[i]{
					count++
				}else{
					count--
				}
			}
			// 对是否过半进行判断
			count =0
			for _,v := range nums{
				if target == v{
					count++
				}
			}
		
			if count > len(nums)/2{
				return target
			}
			return -1
		}
```