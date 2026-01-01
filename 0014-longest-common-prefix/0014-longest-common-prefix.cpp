class Solution {
public:
    string longestCommonPrefix(vector<string>& strs) {
        if (strs.empty()) return ""; // 빈 벡터 예외 처리

        string res = "";
        // 첫 번째 문자열을 기준으로 삼음
        const string& pivot = strs[0]; 

        for (int idx = 0; idx < pivot.size(); ++idx) {
            char cmp = pivot[idx];

            // 나머지 문자열들과 비교
            for (int i = 1; i < strs.size(); ++i) {
                // 1. 현재 인덱스가 비교할 문자열 길이보다 크거나 (범위 초과 방지)
                // 2. 문자가 다르면
                if (idx >= strs[i].size() || strs[i][idx] != cmp) {
                    return res; // 즉시 반환
                }
            }
            res += cmp;
        }
        return res;
    }
};