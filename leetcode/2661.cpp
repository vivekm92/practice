#include <iostream>

using namespace std;

// https://leetcode.com/problems/first-completely-painted-row-or-column/description/
class Solution {
private:
    // T(n) : O(n) ; S(n) : O(n)
    int solvefirstCompleteIndex1(vector<int>& arr, vector<vector<int>>& mat) {

        int sz = arr.size(), m = mat.size(), n = mat[0].size();
        vector<int> brr(sz);
        for (int i = 0; i < sz; i++) {
            brr[arr[i]-1] = i;
        }

        int res = INT_MAX;
        for (int i = 0; i < m; i++) {
            int idx = INT_MIN;
            for (int j = 0; j < n; j++) {
                idx = max(idx, brr[mat[i][j]-1]);
            }
            res = min(res, idx);
        }

        for (int j = 0; j < n; j++) {
            int idx = INT_MIN;
            for (int i = 0; i < m; i++) {
                idx = max(idx, brr[mat[i][j]-1]);
            }
            res = min(res, idx);
        }

        return res;
    }
    // T(n) : O(n) ; S(n) : O(n)
    int solvefirstCompleteIndex(vector<int>& arr, vector<vector<int>>& mat) {
        int sz = arr.size(), n = mat.size(), m = mat[0].size();
        vector<pair<int, int> > brr(sz);
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                brr[mat[i][j]-1] = {i, j};
            }
        }

        unordered_map<int, int> um1, um2;
        for (int i = 0; i < sz; i++) {
            int idx = arr[i] - 1;
            pair<int, int> p = brr[idx];
            int r = p.first, c = p.second;
            um1[r]++;
            um2[c]++;
            if (um1[r] == m || um2[c] == n) {
                return i;
            }
        }

        return -1;
    }
public:
    int firstCompleteIndex(vector<int>& arr, vector<vector<int>>& mat) {
        return solvefirstCompleteIndex1(arr, mat);
    }
};

int main() {
    vector<vector<int> > matrix {{4,3,5}, {1, 2, 6}};
    vector<int> arr {4,1,3,6,2,5};
    Solution *sol = new Solution();
    cout << sol->firstCompleteIndex(arr, matrix) << "\n";
    return 0;
}
