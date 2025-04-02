#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/minimum-sum-of-mountain-triplets-ii/description/
class Solution {
private:
    // Approach 1
    // using suffix array
    // T(n) : O(n) ; S(n) : O(n)
    int solveMinimumSum(vector<int> &nums) {

        int n = nums.size();
        vector<int> rMin(n, INT_MAX);
        for (int i = n-2; i >= 0; i--) {
            rMin[i] = min(rMin[i+1], nums[i+1]);
        }

        int lMin = INT_MAX, minVal = INT_MAX;
        for (int i = 1; i < n - 1; i++) {
            lMin = min(lMin, nums[i-1]);
            if (nums[i] > lMin && nums[i] > rMin[i]) {
                minVal = min(minVal, lMin + nums[i] + rMin[i]);
            }
        }

        return minVal == INT_MAX ? -1 : minVal;
    }
public:
    int minimumSum(vector<int>& nums) {
        return solveMinimumSum(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
