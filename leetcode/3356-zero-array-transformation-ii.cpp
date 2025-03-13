#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/zero-array-transformation-ii/description/
class Solution {
private:
    // Approach 1
    // brute-force approach ( TLE )
    // T(n) : O(n2) ; S(n) : O(1)
    int solveMinZeroArray(vector<int>& nums, vector<vector<int>>& queries) {

        int count = 0;
        bool flag = true;
        for (auto &num : nums) {
            if (num != 0) {
                flag = false;
                break;
            }
        }
        if (flag) return count;
        for (auto &query : queries) {
            int l = query[0], r = query[1], val = query[2];
            for (int i = l; i <= r; i++) {
                nums[i] = max(nums[i] - val, 0);
            }
            count++;
            bool flag = true;
            for (auto &num : nums) {
                if (num != 0) {
                    flag = false;
                    break;
                }
            }

            if (flag) return count;
        }

        return -1;
    }

    bool checkForAllZeros(vector<int> &nums, vector<vector<int> > &queries, int queryLength) {

        int n = nums.size();
        vector<int> diffArr(n+1, 0);
        for (int i = 0; i < queryLength; i++) {
            int start = queries[i][0], end = queries[i][1], val = queries[i][2];
            diffArr[start] += val;
            diffArr[end+1] -= val;
        }

        int sum = 0;
        for (int i = 0; i < n; i++) {
           sum += diffArr[i];
           if (sum < nums[i]) return false;
        }

        return true;
    }

    // Approach 2
    // using binary search with difference Array
    // T(n) : O((n+m)logn) ; S(n) : O(n)
    int solveMinZeroArray1(vector<int>& nums, vector<vector<int> > &queries) {

        int n = queries.size();
        int l = 0, r = n;

        if (!checkForAllZeros(nums, queries, n)) return -1;

        while (l <= r) {
            int mid = l + (r - l) / 2;
            if (checkForAllZeros(nums, queries, mid)) {
                r = mid - 1;
            } else {
                l = mid + 1;
            }
        }

        return l;
    }

    // Approach 3
    // using line-sweep
    // T(n) : O(n+m) ; S(n) : O(n)
    int solveMinZeroArray2(vector<int>& nums, vector<vector<int> > &queries) {

        int n = nums.size(), qn = queries.size();
        vector<int> diffArr(n+1, 0);
        int sum = 0, k = 0;
        for (int i = 0; i < n; i++) {
            while (sum + diffArr[i] < nums[i]) {
                k++;
                if (k > qn) return -1;
                int l = queries[k-1][0], r = queries[k-1][1], val = queries[k-1][2];
                if (r >= i) {
                    diffArr[max(l, i)] += val;
                    diffArr[r+1] -= val;
                }
            }
            sum += diffArr[i];
        }

        return k;
    }
public:
    int minZeroArray(vector<int>& nums, vector<vector<int>>& queries) {
        return solveMinZeroArray2(nums, queries);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
