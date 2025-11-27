class Solution {
public:
    vector<int> exclusiveTime(int n, vector<string>& logs) {
        vector<int> res(n,0);

        stack<int> s;
        int prevTime = 0;
        for(auto msg : logs){
            int firColon = msg.find(':');        
            int secColon = msg.find(':', firColon + 1);

            int id = stoi(msg.substr(0, firColon));
            string type = msg.substr(firColon + 1, secColon - firColon - 1);
            int time = stoi(msg.substr(secColon + 1));

            if(type == "start"){
                if(s.empty()){
                    s.push(id);
                    prevTime = time;
                }
                else{
                    int tempId = s.top();
                    res[tempId] += time - prevTime;
                    s.push(id);
                    prevTime = time;
                }
            }
            if(type == "end"){
                int tempId = s.top();
                s.pop();
                res[tempId] += time - prevTime + 1;
                prevTime = time + 1;
            }
        }

        return res;
    }
};