class Solution {
public:
    vector<int> shuffle(vector<int>& nums, int n) {
        int l = 0, r = n;
        vector<int> res;
        for(int i = 0; i < n; i++){
            res.push_back(nums[l++]);
            res.push_back(nums[r++]);
        }

        return res;
    }
};