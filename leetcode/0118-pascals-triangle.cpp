#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/pascals-triangle/description/
class Solution {
private:
    // Approach 1
    // using recurive approach
    // T(n) : O(n*n) ; S(n) : O(n)
    vector<vector<int>> solveGenerate(int n) {

        if (n == 1) return {{1}};
        else if (n == 2) return {{1}, {1,1}};

        vector<vector<int> > res = solveGenerate(n-1);
        vector<int> row = {1};
        for (size_t i = 1; i < res[n-2].size(); i++) {
            row.push_back(res[n-2][i-1] + res[n-2][i]);
        }
        row.push_back(1);
        res.push_back(row);
        return res;
    }

    // Approach 2
    // using iterative approach
    // T(n) : O(n*n) ; S(n) : O(1)
    vector<vector<int> > solveGenerate1(int n) {

        vector<vector<int> > res;
        if (n < 1) return res;
        res.push_back({1});
        if (n == 1) return res;

        for (int i = 2; i <= n; i++) {
            vector<int> temp = {1};
            for (size_t j = 1; j < res[i-2].size(); j++) {
                temp.push_back(res[i-2][j-1] + res[i-2][j]);
            }
            temp.push_back(1);
            res.push_back(temp);
        }

        return res;
    }
public:
    vector<vector<int>> generate(int numRows) {
        return solveGenerate1(numRows);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
