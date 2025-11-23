class Solution {
public:
    vector<int> getConcatenation(vector<int>& nums) {
        vector<int> newVec(2 * nums.size());
        copy(nums.begin(), nums.end(), newVec.begin());
        copy(nums.begin(), nums.end(), newVec.begin() + nums.size());

        return newVec;
    }
};