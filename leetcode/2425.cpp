#include <iostream>

using namespace std;

// https://leetcode.com/problems/bitwise-xor-of-all-pairings/description/
class Solution {
private:
    // using bit manipulation : ( 0^x = x )
    // T(n) : O(n) ; S(n) : O(1)
    int solvexorAllNums(vector<int>& nums1, vector<int>& nums2) {

        int res = 0;
        int n = nums1.size(), m = nums2.size();
        if (n%2 == 0 && m%2 == 0) return res;
        if (n%2 != 0) {
            for (auto &v : nums2) {
                res ^= v;
            }
        }
        if (m%2 != 0) {
            for (auto &v : nums1) {
                res ^= v;
            }
        }
        return res;
    }
public:
    int xorAllNums(vector<int>& nums1, vector<int>& nums2) {
        return solvexorAllNums(nums1, nums2);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> r1{1, 2};
    vector<int> r2{3, 4};
    cout << sol->xorAllNums(r1, r2) << "\n";
    r2.push_back(5); // 1 ^ 2
    cout << sol->xorAllNums(r1, r2) << "\n";
    r1.push_back(6); // 1 ^ 2 ^ 6 ^ 3 ^ 4 ^ 5 = 7
    cout << sol->xorAllNums(r1, r2) << "\n";
    return 0;
}
