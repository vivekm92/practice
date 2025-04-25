#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-of-interesting-subarrays/description/
class Solution {
private:

    // Approach 1
    // using brute-force
    // T(n) : O(n2) ; S(n) : O(1)
    long long solveCountInterestingSubarrays(vector<int> &nums, int modulo, int k) {

        int n = nums.size();
        long long res = 0;
        for (int i = 0; i < n; i++) {
            int count = 0;
            for (int j = i; j < n; j++) {
                if (nums[j] % modulo == k) count++;
                if (count % modulo == k) res++;
            }
        }

        return res;
    }

    // Approach 2
    // using hash-map
    // T(n) : O(n) ; S(n) : O(n)
    long long solveCountInterestingSubarrays1(vector<int> &nums, int modulo, int k) {

        int n = nums.size();
        long long res = 0, prefix = 0;
        unordered_map<int, int> um;
        um[0] = 1;
        for (int i = 0; i < n; i++) {
            prefix += (nums[i] % modulo == k) ? 1 : 0;
            res += um[(prefix - k + modulo) % modulo];
            um[prefix % modulo]++;
        }

        return res;
    }
public:
    long long countInterestingSubarrays(vector<int>& nums, int modulo, int k) {
        return solveCountInterestingSubarrays1(nums, modulo, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
