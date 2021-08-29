package dp21;

/**
 * @ClassName: rob
 * @Description: TODO
 * @Author: jackey
 * @Create: 2021/8/29 下午9:38
 */
public class rob {
    public int rob1(int[] nums) {
        int n = nums.length;
        if (n == 0) {
            return 0;
        }
        if (n == 1) {
            return nums[0];
        }

        int tag1 = nums[0], tag2 = Math.max(nums[0], nums[1]);

        for (int i = 2; i < n; i++) {
            int tmp = tag2;
            tag2 = Math.max(tag2, tag1 + nums[i]);
            tag1 = tmp;
        }
        return tag2;
    }

    public int robRange(int[] nums, int start, int end) {
        int first = nums[start], second = Math.max(nums[start + 1], nums[start]);
        for (int i = start + 2; i <= end; i++) {
            int tmp = second;
            second = Math.max(second, first + nums[i]);
            first = tmp;
        }
        return second;
    }

    public int rob(int[] nums) {
        int n = nums.length;
        if (n == 1) {
            return nums[0];
        } else if (n == 2) {
            return Math.max(nums[0], nums[1]);
        }
        return Math.max(robRange(nums, 0, n - 2), robRange(nums, 1, n - 1));
    }
}
