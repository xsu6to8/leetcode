class Solution {
public:
    vector<string> buildArray(vector<int>& target, int n) {
        vector<string> res;
        unordered_set<int> s;
        for(auto i : target)
            s.insert(i);

        for(int i = 1; i <= target[target.size() - 1]; i++){
            if(s.find(i) != s.end())
                res.push_back("Push");
            else
            {
                res.push_back("Push");
                res.push_back("Pop");
            }
        }
        return res;
    }
};