class Solution {
public:
    void merge(vector<int>& nums1, int m, vector<int>& nums2, int n) {
        int insIdx = m+n-1;
        m -= 1; n -= 1;

        while(insIdx > -1){
            if(m > -1 && n > -1){
                if(nums1[m] > nums2[n])
                    nums1[insIdx] = nums1[m--];
                else
                    nums1[insIdx] = nums2[n--];
            }
            else if (m == -1)
                nums1[insIdx] = nums2[n--];
            else if (n == -1)
                nums1[insIdx] = nums1[m--];
            
            insIdx--;
        }
    }
};