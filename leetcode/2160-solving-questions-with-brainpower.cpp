#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/solving-questions-with-brainpower/description/
class Solution {
private:
    long long solveMostPointsRec(vector<vector<int> > &questions, int idx) {

        if (idx >= questions.size()) {
            return 0;
        }

        return max(solveMostPointsRec(questions, idx + questions[idx][1] + 1) + questions[idx][0], solveMostPointsRec(questions, idx + 1));
    }

    // Approach 1
    // brute-force solution
    // T(n) : O(2^n) ; S(n) ; O(n)
    long long solveMostPoints(vector<vector<int> > &questions) {
        return solveMostPointsRec(questions, 0);
    }

    // Approach 2
    // using bottom-up approach
    // T(n) : O(n) ; S(n) : O(n)
    long long solveMostPoints1(vector<vector<int> > &questions) {

        int n = questions.size();
        vector<long long> dp(n, 0);
        dp[n-1] = questions[n-1][0];
        for (int i = n-2; i >= 0; i--) {
            dp[i] = questions[i][0];
            int skip = questions[i][1];
            if (i + skip + 1 < n) {
                dp[i] += dp[i+skip+1];
            }
            dp[i] = max(dp[i], dp[i+1]);
        }

        return dp[0];
    }

    long long solveMostPointsRec2(vector<vector<int> > &questions, int idx, vector<long long> &dp) {

        if (idx >= questions.size()) {
            return 0;
        }

        if (dp[idx] != -1) return dp[idx];
        dp[idx] = max(solveMostPointsRec2(questions, idx + 1, dp), solveMostPointsRec2(questions, idx + questions[idx][1] + 1, dp) + questions[idx][0]);
        return dp[idx];
    }

    // Approach 3
    // using top-down approach
    // T(n) : O(n) ; S(n) : O(n)
    long long solveMostPoints2(vector<vector<int> > &questions) {
        int n = questions.size();
        vector<long long> dp(n, -1);
        dp[n-1] = questions[n-1][0];
        solveMostPointsRec2(questions, 0, dp);
        return dp[0];
    }
public:
    long long mostPoints(vector<vector<int>>& questions) {
        return solveMostPoints2(questions);
    }
};


// Driver Code for testing
int main() {

    return 0;
}
