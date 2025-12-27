class Solution {
public:
    int lengthOfLongestSubstring(string s) {
        if (s.size() == 0) 
            return 0;

        int maxLen = 0;
        int l = 0;
        unordered_set<char> us;
        for (int r = 0; r < s.size(); r++) {
            while (us.find(s[r]) != us.end()) {
                us.erase(s[l]);
                l++;
            }

            us.insert(s[r]);
            maxLen = max(maxLen, r - l + 1);
        }
        
        return maxLen;
    }
};