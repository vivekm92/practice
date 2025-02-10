#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/longest-strictly-increasing-or-strictly-decreasing-subarray/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(1)
    int solvelongestMonotonicSubarray(vector<int>& nums) {
        int n = nums.size();
        int incCount = 1, decCount = 1, maxIncCount = 1, maxDecCount = 1;
        for (int i = 1; i < n; i++) {
            if (nums[i] > nums[i-1]) incCount++;
            else incCount = 1;
            maxIncCount = max(maxIncCount, incCount);

            if (nums[i] < nums[i-1]) decCount++;
            else decCount = 1;
            maxDecCount = max(maxDecCount, decCount);
        }
        return maxDecCount > maxIncCount ? maxDecCount : maxIncCount;
    }

    // T(n) : O(n) ; S(n) : O(1)
    int solvelongestMonotonicSubarray1(vector<int>& nums) {
        size_t n=nums.size();
        int incCount=1, decCount=1, maxCount=1;
        for (int i=1; i<n; i++) {
            if (nums[i] > nums[i-1]) incCount++;
            else incCount = 1;
            if (nums[i] < nums[i-1]) decCount++;
            else decCount = 1;
            maxCount = max(maxCount, max(incCount,decCount));
        }
        return maxCount;
    }
public:
    int longestMonotonicSubarray(vector<int>& nums) {
        return solvelongestMonotonicSubarray(nums);
    }
};


// Driver Code for testing
int main() {

    return 0;
}

