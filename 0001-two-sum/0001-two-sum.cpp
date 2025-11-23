class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> memo; 

        for(int i = 0; i < nums.size(); i++) {
            int goal = target - nums[i];

            if (memo.find(goal) != memo.end()) {
                return { memo[goal], i }; 
            }

            memo[nums[i]] = i;
        }
        return {};
    }
};