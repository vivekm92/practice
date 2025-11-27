#include <bits/stdc++.h>

using namespace std;

// https://www.geeksforgeeks.org/problems/triangle-path-sum/1
class Solution {
  private:
    int recurse(int i, int j, vector<vector<int> > &t) {
        if (i == t.size() || j == t[i].size()) {
            return 0;
        }
        return t[i][j] + min(recurse(i+1,j,t), recurse(i+1,j+1,t));
    }

    // Approach 1
    // using recursion
    // T(n) : O(2^(n*n)) ; S(n) : O(n*n)
    int solveMinPathSum(vector<vector<int> > &t) {
        return recurse(0, 0, t);
    }

    int memo(int i, int j, vector<vector<int> > &dp, vector<vector<int> > &t) {
        if (i == t.size() || j == t[i].size()) {
            return 0;
        }
        if (dp[i][j] != -1) return dp[i][j];
        return dp[i][j] = t[i][j] + min(memo(i+1,j,dp,t), memo(i+1,j+1,dp,t));
    }

    // Approach 2
    // using memoization ( top-down approach )
    // T(n) : O(n*n) ; S(n) : O(n*n)
    int solveMinPathSum1(vector<vector<int> > &t) {
        int m = t.size(), n = t[m-1].size();
        vector<vector<int> > dp(m, vector<int>(n, -1));
        return memo(0, 0, dp, t);
    }

    // Approach 3
    // using tabulation ( bottom-up approach )
    // T(n) : O(n*n) ; S(n) : O(n*n)
    int solveMinPathSum2(vector<vector<int> > &t) {
        int n = t.size();
        vector<vector<int> > dp(n);
        for (int i = 0; i < n; i++) {
            dp[i] = vector<int>(t[i].size(), -1);
        }

        // base case
        for (int j = 0; j < t[n-1].size(); j++) {
            dp[n-1][j] = t[n-1][j];
        }

        for (int i = n-2; i >= 0; i--) {
            for (int j = 0; j < t[i].size(); j++) {
                dp[i][j] = t[i][j] + min(dp[i+1][j], dp[i+1][j+1]);
            }
        }

        return dp[0][0];
    }

    // Approach 4
    // using tabulation ( optimised bottom-up approach )
    // T(n) : O(n*n) ; S(n) : O(n)
    int solveMinPathSum3(vector<vector<int> > &t) {

        int m = t.size(), n = t[m-1].size();
        vector<int> dp(n, 0);
        for (int j = 0; j < n; j++) {
            dp[j] = t[m-1][j];
        }

        for (int i = n-2; i >= 0; i--) {
            for (int j = 0; j < t[i].size(); j++) {
                dp[j] = t[i][j] + min(dp[j], dp[j+1]);
            }
        }

        return dp[0];
    }
  public:
    int minPathSum(vector<vector<int>>& triangle) {
        return solveMinPathSum3(triangle);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
