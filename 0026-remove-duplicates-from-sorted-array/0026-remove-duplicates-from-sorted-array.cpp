class Solution {
public:
    int removeDuplicates(vector<int>& nums) {
        unordered_set<int> us;
        int idx = 0;
        for(auto i : nums){
            if(us.find(i) == us.end()){
                us.insert(i);
                nums[idx++] = i;
            }
        }
        return us.size();
    }
};