#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-subarrays-where-max-element-appears-at-least-k-times/description/
class Solution {
private:
    // Approach 1
    // using two-pointers
    // T(n) : O(n) : S(n) : O(1)
    long long solveCountSubarrays(vector<int> &nums, int k) {

        int maxVal = *max_element(nums.begin(), nums.end());
        int n = nums.size(), currCount = 0, j = 0;
        long long count = 0;
        for (int i = 0; i < n; i++) {
            while (j < n && currCount < k) {
                if (nums[j] == maxVal) currCount++;
                j++;
            }
            if (currCount < k) break;
            count += (n - j + 1);
            if (nums[i] == maxVal) currCount--;
        }

        return count;
    }

    // Approach 2
    // using two-pointers
    // T(n) : O(n) ; S(n) : O(1)
    long long solveCountSubarrays1(vector<int> &nums, int k) {

        int maxVal = *max_element(nums.begin(), nums.end());
        long long count = 0;
        int n = nums.size(), currCount = 0, j = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i] == maxVal) currCount++;
            while (currCount == k) {
                if (nums[j] == maxVal) currCount--;
                j++;
            }
            count += j;
        }

        return count;
    }

    // Approach 3
    // using math
    // T(n) : O(n) ; S(n) : O(n)
    long long solveCountSubarrays2(vector<int> &nums, int k) {

        int maxVal = *max_element(nums.begin(), nums.end());
        vector<int> maxIndices;
        long long count = 0;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            if (nums[i] == maxVal) maxIndices.push_back(i);
            int freq = maxIndices.size();
            if (freq >= k) count += maxIndices[freq - k] + 1;
        }

        return count;
    }
public:
    long long countSubarrays(vector<int>& nums, int k) {
        return solveCountSubarrays2(nums, k);
    }
};

// Driver Code for testing
int main() {

    Solution *sol = new Solution();
    vector<int> arr {1,3,2,3,3};
    cout << sol->countSubarrays(arr, 2) << "\n";

    return 0;
}
