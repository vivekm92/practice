#include <iostream>

using namespace std;

// https://leetcode.com/problems/unique-binary-search-trees/description/
class Solution {
private:
    // Approach 1
    // using dp
    // T(n) : O(n2) ; S(n) ; O(n)
    int solveNumTrees(int n) {

        vector<int> dp(n+1, 0);
        dp[0] = 1;
        dp[1] = 1;
        for (int i = 2; i <= n; i++) {
            for (int j = 1; j <= i; j++) {
                dp[i] = dp[j-1] * dp[i-j];
            }
        }
        return dp[n];
    }

    // Approach 2
    // using catalan numbers
    // T(n) : O(n) ; S(n) : O(1)
    int solveNumTrees1(int n) {
        long long c = 1;
        for (int i = 0; i < n; i++) {
            c = c * 2 * (2 * i + 1) / (i + 2);
        }
        return (int )c;
    }
public:
    int numTrees(int n) {
        return solveNumTrees1(n);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
