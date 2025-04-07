#include <iostream>
#include <vector>
#include <numeric>

using namespace std;

// https://leetcode.com/problems/partition-equal-subset-sum/description/
class Solution {
private:
    bool solveCanPartitionHelper(vector<int> &nums, int sum, int total, int idx) {
        if (idx >= nums.size()) return false;
        if (sum == total) return true;

        bool include = solveCanPartitionHelper(nums, sum + nums[idx], total, idx + 1);
        bool exclude = solveCanPartitionHelper(nums, sum, total, idx + 1);

        return include || exclude;
    }

    // Approach 1
    // brute-force approach
    // T(n) : O(2^n) ; S(n) : O(n)
    bool solveCanPartition(vector<int> &nums) {
        int total = 0;
        for (auto &num : nums) total += num;
        if (total % 2 != 0) return false;
        int subSetSum = total / 2;
        return solveCanPartitionHelper(nums, 0, subSetSum, 0);
    }

   bool solveCanPartitionHelper1(vector<int> &nums, int idx, int sum, vector<vector<int> > &dp) {
        if (sum == 0) return true;
        else if (idx >= nums.size() || sum < 0) return false;

        if (dp[idx][sum] != -1) return static_cast<bool>(dp[idx][sum]);
        bool include = solveCanPartitionHelper1(nums, idx+1, sum - nums[idx], dp);
        bool exclude = solveCanPartitionHelper1(nums, idx+1, sum, dp);

        dp[idx][sum] = include || exclude;
        return static_cast<bool>(dp[idx][sum]);
    }

    // Approach 2
    // top-down DP
    // T(n) : O(m*n) ; S(n) : O(m*n)
    bool solveCanPartition1(vector<int> &nums) {
        int total = 0;
        for (auto &num : nums) total += num;
        if (total % 2 != 0) return false;
        int n = nums.size(), subSetSum = total / 2;
        vector<vector<int> > dp(n+1, vector<int>(subSetSum+1, -1));
        return solveCanPartitionHelper1(nums, 0, subSetSum, dp);
    }

    // Approach 3
    // bottom-up DP
    // T(n) : O(m*m) ; S(n) : O(m*n)
    bool solveCanPartition2(vector<int> &nums) {
        int total = 0;
        for (auto &num : nums) total += num;
        int n = nums.size(), subSetSum = total / 2;
        if (total % 2 != 0) return false;

        vector<vector<bool> > dp(n+1, vector<bool>(subSetSum+1, false));
        dp[0][0] = true;
        for (int i = 1; i <= n; i++) {
            int curr = nums[i-1];
            for (int j = 0; j <= subSetSum; j++) {
                if (j < curr) {
                    dp[i][j] = dp[i-1][j];
                } else {
                    dp[i][j] = dp[i-1][j] || dp[i-1][j-curr];
                }
            }
        }

        return dp[n][subSetSum];
    }

    // Approach 4
    // optimised DP
    // T(n) : O(n*m) ; S(n) :O(n)
    bool solveCanPartition3(vector<int> &nums) {

        int total = 0;
        for (auto &num : nums) total += num;
        if (total % 2 != 0) return false;

        int n = nums.size(), subSetSum = total / 2;
        vector<bool> dp(subSetSum+1, false);
        dp[0] = true;
        for (int i = 1; i <= n; i++) {
            int curr = nums[i-1];
            for (int j = subSetSum; j >= curr; j--) {
                dp[j] = dp[j] || dp[j-curr];
            }
        }

        return dp[subSetSum];
    }

    // Approach 5
    // using bit-set
    // T(n) : O(n*m) ; S(n) : O(m)
    bool solveCanPartition4(vector<int> &nums) {

        int total = accumulate(nums.begin(), nums.end(), 0);
        if (total % 2 != 0) return false;
        bitset<10001> bit(1);
        for (auto &num: nums) {
            bit |= (bit << num);
            if (bit[total/2]) return true;
        }
        return static_cast<bool>(bit[total/2]);
    }
public:
    bool canPartition(vector<int>& nums) {
        return solveCanPartition3(nums);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
