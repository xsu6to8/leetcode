class Solution {
public:
    int maxSubArray(vector<int>& nums) {
        if(nums.size() == 1)
            return nums[0];

        int currMax = nums[0];
        for (int i = 1; i < nums.size(); i++){
            int curr = nums[i];
            if (nums[i - 1] > 0){
                nums[i] = nums[i] + nums[i - 1];

                if(nums[i] > currMax)
                    currMax = nums[i];
            }
            else{
                if(nums[i] > currMax)
                    currMax = nums[i];
            }
        }

        return currMax;
    }
};