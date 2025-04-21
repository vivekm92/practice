#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

// https://leetcode.com/problems/maximum-size-subarray-sum-equals-k/description/
class Solution {
private:
    // Approach 1
    // using prefixSum + hash-map
    // T(n) : O(n) ; S(n) : O(n)
    int solveMaxSubArrayLen(vector<int> &nums, int k) {

        int n = nums.size(), maxlen = 0;
        long long prefixSum = 0;
        unordered_map<int, int> um;
        for (int i = 0; i < n; i++) {
            prefixSum += 1LL * nums[i];
            if (prefixSum == k) maxlen = i + 1;

            if (um.find(prefixSum - k) != um.end()) {
                maxlen = max(maxlen, i - um[prefixSum - k]);
            }

            if (um.find(prefixSum) == um.end()) {
                um[prefixSum] = i;
            }
        }

        return maxlen;
    }
public:
    int maxSubArrayLen(vector<int>& nums, int k) {
        return solveMaxSubArrayLen(nums, k);
    }
};

// Driver Code for testing
int main() {

}
