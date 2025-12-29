class Solution {
public:
    double findMedianSortedArrays(vector<int>& nums1, vector<int>& nums2) {
        int m = nums1.size();
        int n = nums2.size();
        
        bool isEven = ((m + n) % 2 == 0);
        int idx = (m + n) / 2;
        
        int l = 0; 
        int r = 0; 
        
        int curr = 0; 
        int prev = 0; 

        for(int i = 0; i <= idx; i++){
            prev = curr; 
            
            if (l < m && (r >= n || nums1[l] < nums2[r])) {
                curr = nums1[l++];
            } else {
                curr = nums2[r++];
            }
        }
        
        if (isEven) {
            return (prev + curr) / 2.0;
        } else {
            return curr;
        }
    }
};