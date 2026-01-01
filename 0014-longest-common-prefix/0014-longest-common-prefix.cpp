class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        string res = "";
        
        bool isSame = true;
        int idx = 0;
        while(isSame){
            char cmp = strs[0][idx];
            for (auto curr : strs){
                if(curr[idx] == NULL || curr[idx] != cmp){
                    isSame = false;
                    break;
                }
            }
            if(isSame)
                res += cmp;
            idx++;    
        }
        return res;
    }
};