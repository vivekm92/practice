#include <iostream>

using namespace std;

// https://leetcode.com/problems/bitwise-and-of-numbers-range/description/
class Solution {
private:
    // T(n) : O(1) ; S(n) : O(1)
    int solverangeBitwiseAnd(int left, int right) {
        int res = 0;
        for (int i = 31; i >= 0; i--) {
            if ((left&(1<<i)) ^ (right&(1<<i))) {
                break;
            }
            if ((left&(1<<i)) & (right&(1<<i))) {
                res += (1<<i);
            }
        }
        return res;
    }
public:
    int rangeBitwiseAnd(int left, int right) {
        return solverangeBitwiseAnd(left, right);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->rangeBitwiseAnd(3, 5) << "\n";
    cout << sol->rangeBitwiseAnd(7, 5) << "\n";
    cout << sol->rangeBitwiseAnd(0, 0) << "\n";
    cout << sol->rangeBitwiseAnd(1, INT_MAX) << "\n";
    return 0;
}
