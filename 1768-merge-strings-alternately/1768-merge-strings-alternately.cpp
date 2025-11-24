class Solution {
public:
    string mergeAlternately(string word1, string word2) {
        string res = "";
        int max_len = max(word1.length(),word2.length());
        int one = 0, two = 0;
        for (int i = 0; i < max_len; i++){
            if(one != word1.length()){
                res += word1[one++];
            }
            if(two != word2.length()){
                res += word2[two++];
            }
        }
        return res;
    }
};