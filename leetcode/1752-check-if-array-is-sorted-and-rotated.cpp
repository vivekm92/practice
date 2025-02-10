#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/check-if-array-is-sorted-and-rotated/description/
class Solution {
private:
    // brute-force approach
    // T(n) : O(n2) ; S(n) : O(1)
    bool solvecheck(vector<int>& nums) {
        int n = nums.size();
        for (int i = 0; i < n; i++) {
            bool flag = true;
            for (int j = i; j < i + n - 1; j++) {
                if (nums[j%n] > nums[(j+1)%n]) {
                    flag = false;
                    break;
                }
            }
            if (flag) return flag;
        }
        return false;
    }

    // T(n) : O(n) ; S(n) : O(1)
    bool solvecheck1(vector<int>& nums) {
        int n = nums.size(), inversionCount = 0;
        for (int i = 0; i < n; i++) {
            if (nums[i%n] > nums[(i+1)%n]) inversionCount++;
        }
        return inversionCount <= 1;
    }
public:
    bool check(vector<int>& nums) {
        return solvecheck1(nums);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    vector<int> arr {1, 2, 3, 4, 5};
    vector<int> arr1 {3, 4, 5, 1, 2};
    vector<int> arr2 {5, 1, 2, 3, 4};
    vector<int> arr3 {1, 3, 2, 4, 5};
    cout << boolalpha;
    cout << sol->check(arr) << "\n";
    cout << sol->check(arr1) << "\n";
    cout << sol->check(arr2) << "\n";
    cout << sol->check(arr3) << "\n";
    return 0;
}
