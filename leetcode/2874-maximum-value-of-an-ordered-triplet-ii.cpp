#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-ii/description/
class Solution {
private:
    // Approach 1
    // using greedy + prefix suffix array
    // T(n) : O(n) ; S(n) : O(n)
    long long solveMaximumTripletValue(vector<int> &nums) {

        int n = nums.size();
        vector<int> lMax(n, 0), rMax(n, 0);
        for (int i = 1; i < n; i++) {
            lMax[i] = max(lMax[i-1], nums[i-1]);
        }

        for (int i = n-2; i >= 0; i--) {
            rMax[i] = max(rMax[i+1], nums[i+1]);
        }

        long long maxVal = 0;
        for (int j = 1; j < n - 1; j++) {
            maxVal = max(maxVal, (long long )(lMax[j] - nums[j]) * rMax[j]);
        }

        return maxVal;
    }

    // Approach 2
    // using optimised greedy
    // T(n) : O(n) ; S(n) : O(1)
    long long solveMaximumTripletValue1(vector<int> &nums) {

        int n = nums.size();
        long long maxVal = 0, maxDiff = 0, maxi = 0;
        for (int i = 0; i < n; i++) {
            maxVal = max(maxVal, maxDiff * nums[i]);
            maxDiff = max(maxDiff, maxi - nums[i]);
            maxi = max(maxi, (long long )nums[i]);
        }

        return maxVal;
    }
public:
    long long maximumTripletValue(vector<int>& nums) {
        return solveMaximumTripletValue1(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
