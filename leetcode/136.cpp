#include <iostream>

using namespace std;

// https://leetcode.com/problems/single-number/description/
class Solution {
private:
    // generic method for solving such questions
    // T(n) : O(n) ; S(n) : O(1)
    int solvesingleNumber1(vector<int>& nums) {
        int res = 0;
        for (int i = 0; i < 32; i++) {
            int bitSum = 0;
            for (auto &v : nums) {
                bitSum += (v>>i)&1;
            }
            int x = bitSum % 2;
            res = res | (x<<i);
        }
        return res;
    }
    // using xor-operations
    // T(n) : O(n) ; S(n) : O(1)
    int solvesingleNumber(vector<int>& nums) {
        int res = 0;
        for (auto &v : nums) res ^= v;
        return res;
    }
public:
    int singleNumber(vector<int>& nums) {
        return solvesingleNumber1(nums);
    }
};

// Driver Code for testing
int main(){
    Solution *sol = new Solution();
    vector<int> temp {1,1,2,2,3};
    cout << sol->singleNumber(temp) << "\n";
    return 0;
}
