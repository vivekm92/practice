#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/maximum-ascending-subarray-sum/description/
class Solution {
private:
    // using two-pointers
    // T(n) : O(n) ; S(n) : O(1)
    int solvemaxAscendingSum(vector<int>& nums) {
        int n=nums.size();
        if (n==0) return 0;
        int right=1, maxSum=nums[0];
        while (right < n) {
            int currSum=nums[right-1];
            while (right<n && nums[right] > nums[right-1]) {
                currSum += nums[right];
                right++;
            }
            maxSum = max(maxSum, currSum);
            right++;
        }
        return maxSum;
    }

    // linear-traversal
    // T(n) : O(n) ; S(n) : O(1)
     int solvemaxAscendingSum1(vector<int>& nums) {
         int n=nums.size();
         if (n==0) return 0;
         int currSum=nums[0], maxSum=nums[0];
         for (int i=1; i<n; i++) {
             if (nums[i] > nums[i-1]) {
                 currSum += nums[i];
             } else {
                 currSum = nums[i];
             }
             maxSum = max(maxSum, currSum);
         }
         return maxSum;
     }
public:
    int maxAscendingSum(vector<int>& nums) {
        return solvemaxAscendingSum1(nums);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> arr {10, 20, 30, 5, 500};
    cout << sol->maxAscendingSum(arr) << "\n";
    return 0;
}
