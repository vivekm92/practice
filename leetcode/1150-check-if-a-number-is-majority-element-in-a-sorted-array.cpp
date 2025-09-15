#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/check-if-a-number-is-majority-element-in-a-sorted-array/description/
class Solution {
private:
    // Approach 1
    // frequency count
    // T(n) : O(n) ; S(n) : O(1)
    bool solveIsMajorityElement(vector<int> &nums, int target) {

        int count = 0;
        for (auto &num : nums) {
            if (num == target) count++;
        }

        return count > (nums.size() / 2);
    }

    // Approach 2
    // using binary-search
    // T(n) : O(logn) ; S(n) : O(1)
    bool solveIsMajorityElement1(vector<int> &nums, int target) {

        int n = nums.size();
        int l = 0, r = n;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] < target) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        int idx1 = l;

        l = 0;
        r = n;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] <= target) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        int idx2 = l;

        int count = idx2 - idx1;
        return count > (n / 2);
    }

    // Approach 3
    // using STL
    // T(n) : O(logn) ; S(n) : O(1)
    bool solveIsMajorityElement2(vector<int> &nums, int target) {

        int n = nums.size();
        int idx1 = lower_bound(nums.begin(), nums.end(), target) - nums.begin();
        int idx2 = upper_bound(nums.begin(), nums.end(), target) - nums.begin();

        int count = idx2 - idx1;
        return count > (n / 2);
    }

    // Approach 4
    // using binary-search ( optimised )
    // T(n) ; O(logn) ; S(n) : O(1)
    bool solveIsMajorityElement3(vector<int> &nums, int target) {

        int n = nums.size();
        int l = 0, r = n;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] < target) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }

        if ((l + n/2 < n) && (nums[l+n/2] == target)) {
            return true;
        }

        return false;
    }
public:
    bool isMajorityElement(vector<int>& nums, int target) {
        return solveIsMajorityElement(nums, target);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
