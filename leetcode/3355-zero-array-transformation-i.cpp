#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/zero-array-transformation-i/description/
class Solution {
private:
    // Approach 1
    // using difference array
    // T(n) : O(n) ; S(n) : O(n)
    bool solveIsZeroArray(vector<int> &nums, vector<vector<int> > &queries) {

        int n = nums.size();
        vector<int> arr(n+1, 0);
        for (auto &query : queries) {
            arr[query[0]]--;
            arr[query[1]+1]++;
        }

        for (int i = 0; i < n; i++) {
            if (i > 0) arr[i] += arr[i-1];
            if (nums[i] + arr[i] > 0) return false;
        }

        return true;
    }
public:
    bool isZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        return solveIsZeroArray(nums, queries);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
