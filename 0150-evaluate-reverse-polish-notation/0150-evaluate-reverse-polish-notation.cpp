class Solution {
public:
    int evalRPN(vector<string>& tokens) {
        stack<int> st;
        
        for(string& token : tokens) {
            if (token == "+" || token == "-" || token == "*" || token == "/") {
                int t2 = st.top(); 
                st.pop(); 
                int t1 = st.top(); 
                st.pop(); 
                
                if (token == "+") 
                    st.push(t1 + t2);
                else if (token == "-") 
                    st.push(t1 - t2);
                else if (token == "*") 
                    st.push(t1 * t2);
                else if (token == "/")  
                    st.push(t1 / t2);
            } 
            else {
                st.push(stoi(token));
            }
        }
        return st.top();
    }
};