#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/count-the-number-of-fair-pairs/description/
class Solution {
private:
    int bsearch_lower_bound(vector<int> &nums, int l, int r, int target) {

        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] < target) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        return l;
    }

    int bsearch_lower_bound_1(vector<int> &nums, int l, int r, int target) {

        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] <= target) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        return l;
    }

    // Approach 1
    // Using binary-search
    // T(n) : O(nlogn) ; S(n) : O(logn)
    long long solveCountFairPairs(vector<int> &nums, int lower, int upper) {

        sort(nums.begin(), nums.end());
        long long count = 0;
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            int idx1 = bsearch_lower_bound(nums, i + 1, n, lower - nums[i]);
            int idx2 = bsearch_lower_bound_1(nums, i + 1, n, upper - nums[i]);
            count += (idx2 - idx1);
        }

        return count;
    }

    // Approach 2
    // using binary - search
    // T(n) : O(nlogn) ; S(n) : O(logn)
    long long solveCountFairPairs1(vector<int> &nums, int lower, int upper) {

        sort(nums.begin(), nums.end());
        int n = nums.size();
        long long count = 0;
        for (int i = 0; i < n; i++) {
            int idx1 = bsearch_lower_bound(nums, i+1, n, lower - nums[i]);
            int idx2 = bsearch_lower_bound(nums, i+1, n, upper - nums[i] + 1);
            count += (idx2 - idx1);
        }

        return count;
    }

    long long bsearch_lower_bound_2(vector<int> &nums, int target) {

        int n = nums.size();
        int l = 0, r = n-1;
        long long res = 0;
        while (l < r) {
            int sum = nums[l] + nums[r];
            if (sum < target) {
                res += (r - l);
                l++;
            } else {
                r--;
            }
        }
        return res;
    }

    // Approach 3
    // using two pointers
    // T(n) : O(nlogn) ; S(n) : O(logn)
    long long solveCountFairPairs2(vector<int> &nums, int lower, int upper) {

        sort(nums.begin(), nums.end());
        return bsearch_lower_bound_2(nums, upper + 1) - bsearch_lower_bound_2(nums, lower);
    }
public:
    long long countFairPairs(vector<int>& nums, int lower, int upper) {
        return solveCountFairPairs2(nums, lower, upper);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
