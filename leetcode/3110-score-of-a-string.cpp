#include <iostream>

using namespace std;

// https://leetcode.com/problems/score-of-a-string/description/
class Solution {
private:
    // Approach 1
    // using basic simulation
    // T(n) : O(n) ; S(n) : O(1)
    int solveScoreOfString(string s) {

        int n = s.size(), score = 0;
        for (int i = 1; i < n; i++) {
            score += abs(s[i] - s[i-1]);
        }

        return score;
    }
public:
    int scoreOfString(string s) {
        return solveScoreOfString(s);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
