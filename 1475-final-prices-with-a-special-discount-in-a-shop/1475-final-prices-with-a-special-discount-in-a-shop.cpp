class Solution {
public:
    vector<int> finalPrices(vector<int>& prices) {
        stack<int> st;
        st.push(prices[prices.size() -1]);

        vector<int> minus(prices.size(), 0);
        for(int i = prices.size() - 2; i >= 0; i--){
            //  내 뒤에 나오는 얘가 나보다 더 큼
            while (!st.empty() && st.top() > prices[i]){
                st.pop();
            }

            if(!st.empty())
                minus[i] = st.top();

            st.push(prices[i]);   
        }

        vector<int> res;
        for(int i = 0; i < prices.size(); i++){
            res.push_back(prices[i] - minus[i]);
        }

        return res;
    }
};