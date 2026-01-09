class Solution {
public:
    int singleNumber(vector<int>& nums) 
    {
        sort(nums.begin(), nums.end());

        for(auto i : nums)
        {
            int gap = upper_bound(nums.begin(), nums.end(), i) - lower_bound(nums.begin(), nums.end(), i);

            if(gap == 1)
            {
                return i;
            }
        }
        return 0;
    }
};