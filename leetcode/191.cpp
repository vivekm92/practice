#include <iostream>

using namespace std;

// https://leetcode.com/problems/number-of-1-bits/description/
class Solution {
private:
    // T(n) : O(logn) ; S(n) : O(1)
    int solvehammingWeight(int n) {
        int count = 0;
        while (n > 0) {
            count += (n & 1);
            n = n >> 1;
        }
        return count;
    }
public:
    int hammingWeight(int n) {
        return solvehammingWeight(n);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->hammingWeight(10) << "\n";
    cout << sol->hammingWeight(26) << "\n";
    cout << sol->hammingWeight(0) << "\n";
    cout << sol->hammingWeight(7) << "\n";
}
