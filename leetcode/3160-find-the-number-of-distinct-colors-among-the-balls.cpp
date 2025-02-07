#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/find-the-number-of-distinct-colors-among-the-balls/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solvequeryResults(int limit, vector<vector<int>>& queries) {
        
        vector<int> res;
        unordered_map<int, int> ballColors, distinctColors;
        for (auto &query : queries) {
            int x = query[0], y = query[1];
            if (ballColors.find(x) != ballColors.end()) {
                int color = ballColors[x];
                distinctColors[color]--;
                if (distinctColors[color] == 0) distinctColors.erase(color);
            }
            ballColors[x] = y;
            distinctColors[y]++;
            res.push_back(distinctColors.size());
        }
        return res;
    }
public:
    vector<int> queryResults(int limit, vector<vector<int>>& queries) {
        return solvequeryResults(limit, queries);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
