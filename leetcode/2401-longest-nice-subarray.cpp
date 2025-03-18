#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/longest-nice-subarray/description/
class Solution {
private:
    // using sliding-window
    // T(n) : O(n) ; S(n) : O(1)
    int solveLongestNiceSubarray(vector<int>& nums) {

        int maxLength = 0, useBits = 0;
        int start = 0, n = nums.size();
        for (int end = 0; end < n; end++) {
            while ((useBits & nums[end]) != 0) {
                useBits ^= nums[start];
                start++;
            }
            useBits |= nums[end];
            maxLength = max(maxLength, end - start + 1);
        }

        return maxLength;
    }
public:
    int longestNiceSubarray(vector<int>& nums) {
        return solveLongestNiceSubarray(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
