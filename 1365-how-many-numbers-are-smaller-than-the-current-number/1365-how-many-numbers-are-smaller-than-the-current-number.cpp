class Solution {
public:
    vector<int> smallerNumbersThanCurrent(vector<int>& nums) {
        vector<int> sv = nums;
        sort(sv.begin(), sv.end());

        vector<int> res;
        for(auto i : nums)
            res.push_back(lower_bound(sv.begin(), sv.end(), i) - sv.begin());
        
        return res;
    }
};