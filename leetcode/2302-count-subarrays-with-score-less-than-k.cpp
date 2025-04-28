#include <iostream>

using namespace std;

// https://leetcode.com/problems/count-subarrays-with-score-less-than-k/description/
class Solution {
private:
    // Approach 1
    // brute-force
    // T(n) : O(n2) ; S(n) : O(1)
    long long solveCountSubarrays(vector<int> &nums, long long k) {

        int n = nums.size();
        long long count = 0;
        for (int i = 0; i < n; i++) {
            long long currSum = 0;
            for (int j = i; j < n; j++) {
                currSum += nums[j];
                if (currSum * (j-i+1) < k) count++;
            }
        }

        return count;
    }

    // Approach 2
    // using two-pointers
    // T(n) : O(n) ; S(n) : O(1)
    long long solveCountSubarrays1(vector<int> &nums, long long k) {

        int n = nums.size(), idx = 0;
        long long count = 0, currSum = 0;
        for (int i = 0; i < n; i++) {
            currSum += nums[i];
            while (idx <= i && currSum * (i-idx+1) >= k) {
                currSum -= nums[idx];
                idx++;
            }
            count += (i-idx+1);
        }

        return count;
    }
public:
    long long countSubarrays(vector<int>& nums, long long k) {
        return solveCountSubarrays1(nums, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
