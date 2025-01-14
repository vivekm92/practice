#include <iostream>
#include <vector>
#include <numeric>
#define ll long long

using namespace std;

// https://leetcode.com/problems/number-of-ways-to-split-array/description/
class Solution {
    private:
        // using two pass method
        // T(n) : O(n) ; S(n) : O(1)
        int solvewaysToSplitArray(vector<int>& nums) {
            ll totalSum = accumulate(nums.begin(), nums.end(), 0ll), currSum = 0;
            int totalWays = 0, n = nums.size();
            for (int i = 0; i < n - 1; i++) {
                currSum += nums[i];
                totalSum -= nums[i];
                if (currSum >= totalSum) totalWays++;
            }

            return totalWays;
        }
    public:
        int waysToSplitArray(vector<int>& nums) {
            return solvewaysToSplitArray(nums);
        }
};

// Driver Code for testing
int main() {

    Solution *sol = new Solution();
    vector<int> arr = {7,6,5,4,3,2,1};
    cout << sol->waysToSplitArray(arr) << "\n";

    return 0;
}
