
#include<vector>
#include<algorithm>

class Solution{
    public:
        int maxSubAttay(vector<int>& nums){
            int pre = 0,maxAns = nums[0];
            for (auto &&x : nums)
            {
                pre = max(pre+x,x);
                maxAns = max(maxAns,pre);
            }
            return maxAns;
        }
}