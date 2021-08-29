package dp21;

/**
 * @ClassName: maxSubArray
 * @Description: dp 03
 * @Author: jackey
 * @Create: 2021/8/29 下午9:20
 */
class Solution {
    public int maxSubArray(int[] nums) {
        int max = nums[0], curSum = 0;
        for (int x : nums) {
            curSum = Math.max(curSum + x, x);
            max = Math.max(curSum, max);
        }
        return max;
    }
}
