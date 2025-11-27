#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/largest-unique-number/description/
class Solution {
private:
    // Approach 1
    // T(n) : O(n) ; S(n) : O(n)
    int solveLargestUniqueNumber(vector<int> &nums) {

        unordered_map<int, int> um;
        for (auto &n : nums) um[n]++;

        int res = -1;
        for (auto it = um.begin(); it != um.end(); it++) {
            if (it->second == 1) res = max(res, it->first);
        }

        return res;
    }
public:
    int largestUniqueNumber(vector<int>& nums) {
        return solveLargestUniqueNumber(nums);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
