#include <iostream>

using namespace std;

// https://leetcode.com/problems/single-number-iii/description/
class Solution {
    // T(n) : O(n) ; S(n) : O(1)
    vector<int> solvesingleNumber(vector<int>& nums) {
        long long a_b = 0;
        for (auto &num : nums) a_b ^= num;
        int diff = a_b & (-a_b), x = 0;
        for (auto &num : nums) {
            if ((diff&num) != 0) x ^= num;
        }
        return {x, x^(int )a_b};
    }
public:
    vector<int> singleNumber(vector<int>& nums) {
        return solvesingleNumber(nums);
    }
};

// Driver Code for testing
int main() {
    vector<int> temp = {1,1,2,2,3,4};
    Solution *sol = new Solution();
    vector<int> res = sol->singleNumber(temp);
    for (auto &v : res) cout << v << " ";
    cout << "\n";
    return 0;
}
