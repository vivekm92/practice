#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/maximum-absolute-sum-of-any-subarray/description/
class Solution {
private:
    // using kadane's Algorithm
    // T(n) : O(n) ; S(n) : O(1)
    int solveMaxAbsoluteSum(vector<int>& nums) {
        
        int maxSum = INT_MIN, currMaxSum = 0;
        int minSum = INT_MAX, currMinSum = 0, maxAbsSum = 0;
        for (auto &num : nums) {
            currMaxSum = max(currMaxSum + num, num);
            maxSum = max(maxSum, currMaxSum);
            currMinSum = min(currMinSum + num, num);
            minSum = min(minSum, currMinSum);
            maxAbsSum = max(abs(minSum), maxSum);
        }

        return maxAbsSum;
    }
public:
    int maxAbsoluteSum(vector<int>& nums) {
        return solveMaxAbsoluteSum(nums);
    }
};

// Driver Code for testing
int main() {
    vector<int> arr = {1,-3,2,3,-4};        // o/p -> 5
    Solution *sol = new Solution();
    cout << sol->maxAbsoluteSum(arr) << "\n";
    
    vector<int> arr1 = {1,-3,2,3,-4,-3};    // o/p -> 7
    cout << sol->maxAbsoluteSum(arr1) << "\n";
    return 0;
}
