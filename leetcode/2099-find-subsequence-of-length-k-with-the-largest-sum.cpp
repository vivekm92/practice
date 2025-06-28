#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/find-subsequence-of-length-k-with-the-largest-sum/description/
class Solution {
private:
    // Approach 1
    // using sort
    // T(n) : O(nlogn) ; S(n) : O(n)
    vector<int> solveMaxSubsequence(vector<int> &nums, int k) {

        int n = nums.size();
        vector<pair<int, int> > arr;
        for (int i = 0; i < n; i++) {
            arr.push_back({nums[i], i});
        }

        sort(arr.begin(), arr.end(), [&](pair<int, int> &p1, pair<int, int> &p2) {
            return p1.first > p2.first;
        });

        sort(arr.begin(), arr.begin() + k, [&](pair<int, int> &p1, pair<int, int> &p2) {
            return p1.second < p2.second;
        });

        vector<int> res;
        for (int i = 0; i < k; i++) {
            res.push_back(arr[i].first);
        }

        return res;
    }
public:
    vector<int> maxSubsequence(vector<int>& nums, int k) {
        return solveMaxSubsequence(nums, k);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
