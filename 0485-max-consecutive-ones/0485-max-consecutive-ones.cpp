class Solution {
public:
    int findMaxConsecutiveOnes(vector<int>& nums) {
        int max_cnt = 0; 
        int cur_cnt = 0; 

        for(int n : nums) {
            if(n == 1) {
                cur_cnt++; 
                max_cnt = max(max_cnt, cur_cnt); 
            } else {
                cur_cnt = 0; 
            }
        }
        return max_cnt;
    }
};