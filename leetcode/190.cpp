#include <iostream>

using namespace std;

// https://leetcode.com/problems/reverse-bits/description/
class Solution {
private:
    uint32_t solvereverseBits(uint32_t n) {
        uint32_t r = 0;
        for (int i = 0; i < 32; i++) {
           r = r * 2 + ((n >> i) & 1);
        }
        return r;
    }
public:
    // T(n) : O(logn) ; S(n) : O(1)
    uint32_t reverseBits(uint32_t n) {
        return solvereverseBits(n);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->reverseBits(5) << "\n";
}
