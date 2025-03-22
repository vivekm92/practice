#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/description/
class Solution {
private:
    // Approach 1
    // using sliding-window
    // T(n) : O(n) ; S(n) : O(1)
    int solveMinOperations(vector<int> &nums) {

        int n = nums.size(), count = 0;
        for (int i = 0; i < n-2; i++) {
            if (nums[i] == 0) {
                for (int j = i; j < i + 3; j++) {
                    nums[j] = nums[j] ^ 1;
                }
                count++;
            }
        }

        for (int i = n-2; i < n; i++) {
            if (nums[i] == 0) return -1;
        }

        return count;
    }
public:
    int minOperations(vector<int>& nums) {
        return solveMinOperations(nums);
    }
};

// Driver Code for testing
int main() {

    Solution *sol = new Solution();
    vector<int> arr = {1, 0, 1, 1, 0};
    vector<int> arr1 = {0, 1, 0, 0};
    cout << sol->minOperations(arr) << "\n";
    cout << sol->minOperations(arr1) << "\n";
    return 0;
}
