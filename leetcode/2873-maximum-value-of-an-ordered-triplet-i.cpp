#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/maximum-value-of-an-ordered-triplet-i/description/
class Solution {
private:
    // Approach 1
    // brute-force approach
    // T(n) : O(n3) ; S(n) : O(1)
    long long solveMaximumTripletValue(vector<int> &nums) {

        int n = nums.size();
        long long maxVal = 0;
        for (int i = 0; i < n; i++) {
            for (int j = i + 1; j < n; j++) {
                for (int k = j + 1; k < n; k++) {
                    maxVal = max(maxVal, (long long )(nums[i] - nums[j]) * nums[k]);
                }
            }
        }

        return maxVal;
    }

    // Approach 2
    // using greedy approach
    // T(n) : O(n2) ; S(n) : O(1)
    long long solveMaximumTripletValue1(vector<int> &nums) {

        int n = nums.size();
        if (n == 0) return 0;
        long long maxElement = 0, m = max(nums[0], 0);
        for (int j = 1; j < n; j++) {
            for (int k = j + 1; k < n; k++) {
                maxElement = max(maxElement, (long long)(m - nums[j]) * nums[k]);
            }
            m = max(m, (long long)nums[j]);
        }

        return maxElement;
    }

    // Approach 3
    // using Greedy + prefix suffix array
    // T(n) : O(n) ; S(n) : O(n)
    long long solveMaximumTripletValue2(vector<int> &nums) {

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

    // Approach 4
    // using optimised Greedy
    // T(n) : O(n) ; S(n) : O(1)
    long long solveMaximumTripletValue3(vector<int> &nums) {

        int n = nums.size();
        long long maxDiff = 0, maxi = 0, maxVal = 0;
        for (int k = 0; k < n; k++) {
            maxVal = max(maxVal, maxDiff * nums[k]);
            maxDiff = max(maxDiff, maxi - nums[k]);
            maxi = max(maxi, (long long )nums[k]);
        }

        return maxVal;
    }
public:
    long long maximumTripletValue(vector<int>& nums) {
        return solveMaximumTripletValue3(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
