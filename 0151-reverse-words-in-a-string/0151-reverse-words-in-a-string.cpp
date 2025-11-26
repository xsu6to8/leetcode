class Solution {
public:
    string reverseWords(string s) {
        string res = "";
        string tmp = "";
        
        for (auto c : s) {
            if (c == ' ') {
                if (tmp == "") 
                    continue;

                if (res != "") 
                    res = tmp + " " + res;
                else 
                    res = tmp;
                tmp = "";
            } 
            else {
                tmp += c;
            }
        }
        
        if (tmp != "") {
            if (res != "") 
                res = tmp + " " + res;
            else 
                res = tmp;
        }
        
        return res;
    }
};