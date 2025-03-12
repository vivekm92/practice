#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/maximum-count-of-positive-integer-and-negative-integer/description/
class Solution {
private:
    // using brute-force approach
    // T(n) : O(n) ; S(n) : O(1)
    int solveMaximumCount(vector<int> &nums) {

        int pCount = 0 , nCount = 0;
        for (auto &num : nums) {
            if (num > 0) pCount++;
            else if (num < 0) nCount++;
        }

        return pCount > nCount ? pCount : nCount;
    }

    // using binary-search
    // T(n) ; O(logN) ; S(n) : O(1)
    int solveMaximumCount1(vector<int> &nums) {

        int n = nums.size();
        int l = 0, r = n;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] <= 0) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }

        int pCount = n - l;
        l = 0;
        r = n;
        while (l < r) {
            int mid = l + (r - l) / 2;
            if (nums[mid] < 0) {
                l = mid + 1;
            } else {
                r = mid;
            }
        }
        int nCount = l;

        return pCount > nCount ? pCount : nCount;
    }
public:
    int maximumCount(vector<int>& nums) {
        return solveMaximumCount1(nums);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> arr = {-3,-2,-1,0,1,2,3};
    vector<int> arr1 = {-3,-2,-1,0,0,1,2,3};
    vector<int> arr2 = {0,0,1,2,3};
    vector<int> arr3 = {1,2,3};
    vector<int> arr4 = {5,10,15,20};
    vector<int> arr5 = {-1};

    cout << sol->maximumCount(arr) << "\n";
    cout << sol->maximumCount(arr1) << "\n";
    cout << sol->maximumCount(arr2) << "\n";
    cout << sol->maximumCount(arr3) << "\n";
    cout << sol->maximumCount(arr4) << "\n";
    cout << sol->maximumCount(arr5) << "\n";

    return 0;
}
