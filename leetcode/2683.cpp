#include <iostream>

using namespace std;

// https://leetcode.com/problems/neighboring-bitwise-xor/description/
class Solution {
private:
    // using bit-manipulation
    // T(n) : O(n) ; S(n) : O(1)
    bool solvedoesValidArrayExist(vector<int>& derived) {
        int res = 0;
        for (auto &v : derived) {
            res ^= v;
        }
        return res == 0;
    }
public:
    bool doesValidArrayExist(vector<int>& derived) {
        return solvedoesValidArrayExist(derived);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
