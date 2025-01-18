#include <iostream>

using namespace std;

// https://leetcode.com/problems/single-number-ii/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(1)
    int solvesingleNumber(vector<int>& nums) {
        int res = 0;
        for (int i = 0; i < 32; i++) {
            int bitSum = 0;
            for (auto &v : nums) {
                bitSum += (v>>i)&1;
            }
            int x = bitSum % 3;
            res = res | (x<<i);
        }
        return res;
    }
public:
    int singleNumber(vector<int>& nums) {
        return solvesingleNumber(nums);
    }
};


// Driver Code for testing
int main() {
    vector<int> temp {1,1,1,2,2,2,3};
    Solution *sol = new Solution();
    cout << sol->singleNumber(temp) << "\n";
    return 0;
}
