#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/count-square-submatrices-with-all-ones/description/
class Solution {
private:
    // Approach 1
    // bottom-up dp
    // T(n) : O(m*n) ; S(n) : O(m*n)
    int solveCountSquares(vector<vector<int> > &matrix) {

        int count = 0;
        int m = matrix.size(), n = matrix[0].size();
        vector<vector<int> > dp(m+1, vector<int>(n+1, 0));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 1) {
                    dp[i+1][j+1] = min({dp[i+1][j], dp[i][j+1], dp[i][j]}) + 1;
                    count += dp[i+1][j+1];
                }
            }
        }

        return count;
    }

    int recurse(int i, int j, vector<vector<int> > &matrix, vector<vector<int> > &dp) {

        if (i >= (int )matrix.size() || j >= (int )matrix[0].size()) return 0;
        if (matrix[i][j] == 0) return 0;

        if (dp[i][j] != -1) return dp[i][j];

        int right = recurse(i, j+1, matrix, dp);
        int left = recurse(i+1, j, matrix, dp);
        int prev = recurse(i+1, j+1, matrix, dp);

        return dp[i][j] = 1 + min(right, min(left, prev));
    }

    // Approach 2
    // top-down dp
    // T(n) : O(m*n) ; S(n) : O(m*n)
    int solveCountSquares1(vector<vector<int> > &matrix) {

        int count = 0;
        int m = matrix.size(), n = matrix[0].size();
        vector<vector<int> > dp(m, vector<int>(n, -1));
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                count += recurse(i, j, matrix, dp);
            }
        }

        return count;
    }

    // Approach 3
    // optimised bottom-up dp
    // T(n) : O(m*n) ; S(n) : O(n)
    int solveCountSquares2(vector<vector<int> > &matrix) {

        int count = 0, prev = 0;
        int m = matrix.size(), n = matrix[0].size();
        vector<int> dp(n+1, 0);
        for (int i = 1; i <= m; i++) {
            for (int j = 1; j <= n; j++) {
                if (matrix[i-1][j-1] == 1) {
                    int t = dp[j];
                    dp[j] = 1 + min(prev, min(dp[j], dp[j-1]));
                    prev = t;
                    count += dp[j];
                } else {
                    dp[j] = 0;
                }
            }
        }
        return count;
    }
public:
    int countSquares(vector<vector<int>>& matrix) {
        return solveCountSquares2(matrix);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
