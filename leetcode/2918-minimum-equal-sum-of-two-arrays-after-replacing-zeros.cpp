#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/minimum-equal-sum-of-two-arrays-after-replacing-zeros/description/
class Solution {
private:
    // Approach 1
    // using greedy approach
    // T(n) ; O(n) ; S(n) : O(1)
    long long solveMinSum(vector<int> &nums1, vector<int> &nums2) {
        
        int bs1 = 0, bs2 = 0;
        long long sum1 = 0, sum2 = 0;
        for (auto &num : nums1) {
            if (num == 0) bs1++;
            sum1 += num;
        }
        sum1 += bs1;
        
        for (auto &num : nums2) {
            if (num == 0) bs2++;
            sum2 += num;
        }
        sum2 += bs2;
        
        if (sum1 < sum2 && bs1 == 0 || sum2 < sum1 && bs2 == 0) return -1;
        return max(sum1, sum2);
    }
public:
    long long minSum(vector<int>& nums1, vector<int>& nums2) {
        return solveMinSum(nums1, nums2);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}

