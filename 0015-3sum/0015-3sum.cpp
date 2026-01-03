class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> res;
        sort(nums.begin(), nums.end());
        
        for(int i = 0; i < nums.size(); i++) {
            // i 중복 방지: 이전 값과 같으면 건너뜀
            if (i > 0 && nums[i] == nums[i-1]) 
                continue;
            
            int left = i + 1;
            int right = nums.size() - 1;
            
            while (left < right) {
                int sum = nums[i] + nums[left] + nums[right];
                
                if (sum < 0) {
                    left++;
                } else if (sum > 0) {
                    right--;
                } else {
                    // 합이 0인 경우 결과 저장
                    res.push_back({nums[i], nums[left], nums[right]});
                    
                    // left, right 중복 방지 (다음 다른 숫자가 나올 때까지 이동)
                    while (left < right && nums[left] == nums[left+1]) 
                        left++;
                    while (left < right && nums[right] == nums[right-1]) 
                        right--;
                    
                    left++;
                    right--;
                }
            }
        }
        return res;
    }
};