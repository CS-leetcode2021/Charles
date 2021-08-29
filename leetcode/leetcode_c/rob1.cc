#include <vector>

class Solution
{
public:
    int rob(vector<int> &nums)
    {
        if (nums.empty())
        {
            return 0;
        }
        int size = nums.size();

        if (size == 1)
        {
            return nums[0];
        }

        vector<int> dp = vector<int>(size, 0);
        dp[0] = nums[0];
        dp[1] = max(nums[0], nums[1]);

        for (int i = 2; i < size; i++)
        {
            dp[i] = max(dp[i - 2] + nums[i], dp[i - 1]);
        }
        return dp[size - 1];
    }
public
    int robRange(vector<int> &nums, int start, int end)
    {
        int first = nums[start], second = max(nums[start], nums[start + 1]);
        for (int i = start + 2; i <= end; i++)
        {
            int tmp = second;
            second = max(second, first + nums[i]);
            first = tmp;
        }
        return second;
    }

    int rob2(vector<int> &nums)
    {
        if (nums.empty())
        {
            return 0;
        }
        int size = nums.size();
        if (size == 1)
        {
            return nums[0];
        }else if (size == 2)
        {
            return max(nums[0],nums[1]);
        }
        
        return max(robRange(nums, 0, size - 2), robRange(nums, 1, size - 1));
    }
}