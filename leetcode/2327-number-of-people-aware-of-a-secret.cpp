#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/number-of-people-aware-of-a-secret/description/
class Solution {
private:
    int solveRecursive(int curr, const int &n, const int &delay, const int &forget, vector<int> &dp, const int &MOD) {
        if (dp[curr] != -1) return dp[curr];

        int total = (curr + forget - 1 >= n);
        for (int i = delay; i < forget; i++) {
            if (curr + i > n) break;
            total = (total + solveRecursive(curr+i, n, delay, forget, dp, MOD)) % MOD;
        }

        return dp[curr] = total;
    }

    // Approach 1
    // using top-down approach
    // T(n) : O(n*n) ; S(n) : O(n)
    int solvePeopleAwareOfSecret(int n, int delay, int forget) {

        static const int MOD = 1e9 + 7;
        vector<int> dp(n+1, -1);
        return solveRecursive(1, n, delay, forget, dp, MOD);
    }

    // Approach 2
    // using botton-up approach
    // T(n) : O(n) ; S(n) ; O(n)
    int solvePeopleAwareOfSecret1(int n, int delay, int forget) {

        static const int MOD = 1e9 + 7;
        vector<int> dp(n+1, 0);
        dp[1] = 1;
        int window_count = 0;
        for (int i = 2; i <= n; i++) {
            if (i - delay > 0) {
                window_count = (window_count + dp[i-delay]) % MOD;
            }
            if (i - forget > 0) {
                window_count = (window_count - dp[i-forget] + MOD) % MOD;
            }
            dp[i] = window_count;
        }

        int count = 0;
        for (int i = n-forget+1; i<=n; i++) {
            count = (count + dp[i]) % MOD;
        }

        return count;
    }
public:
    int peopleAwareOfSecret(int n, int delay, int forget) {
        return solvePeopleAwareOfSecret(n, delay, forget);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
