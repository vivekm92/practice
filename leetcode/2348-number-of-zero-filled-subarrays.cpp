#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/number-of-zero-filled-subarrays/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(1)
    long long solveZeroFilledSubarray(vector<int> &nums) {

        int i = 0, n = nums.size();
        long long count = 0;
        while (i < n) {
            int j = i;
            while (j < n && nums[j] == 0) {
                j++;
            }
            long long v = j - i;
            count += (v * (v + 1) / 2);
            i = j + 1;
        }

        return count;
    }
public:
    long long zeroFilledSubarray(vector<int>& nums) {
        return solveZeroFilledSubarray(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
