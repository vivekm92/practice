#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-subarrays-with-fixed-bounds/description/
class Solution {
    // Approach 1
    // using brute-force
    // T(n) : O(n2) ; S(n) : O(1)
    long long solveCountSubarrays(vector<int> &nums, int minK, int maxK) {

        long long count = 0;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            int maxVal = nums[i], minVal = nums[i];
            for (int j = i; j < n; j++) {
                maxVal = max(maxVal, nums[j]);
                minVal = min(minVal, nums[j]);
                if (minVal == minK && maxVal == maxK) count++;
            }
        }

        return count;
    }

    // Approach 2
    // using two-pointers
    // T(n) : O)(n) ; S(n) : O(1)
    long long solveCountSubarrays1(vector<int> &nums, int minK, int maxK) {

        int n = nums.size();
        long long count = 0;
        int left = -1, maxPos = -1, minPos = -1;
        for (int i = 0; i < n; i++) {
            if (nums[i] < minK || nums[i] > maxK) {
                left = i;
            }

            if (nums[i] == minK) {
                minPos = i;
            }
            if (nums[i] == maxK) {
                maxPos = i;
            }

            count += max(0, min(maxPos, minPos) - left);
        }

        return count;
    }
public:
    long long countSubarrays(vector<int>& nums, int minK, int maxK) {
        return solveCountSubarrays1(nums, minK, maxK);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
