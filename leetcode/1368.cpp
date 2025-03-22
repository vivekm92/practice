#include <iostream>

using namespace std;

// https://leetcode.com/problems/minimum-cost-to-make-at-least-one-valid-path-in-a-grid/description/
class Solution {
private:
    // using dynamic-programming
    // T(n) : O(n2) ; S(n) : O(n)
    int solveminCost(vector<vector<int>>& grid) {
        int nR = grid.size(), nC = grid[0].size();
        vector<vector<int> > minChanges(nR, vector<int>(nC, INT_MAX));
        minChanges[0][0] = 0;
        while (true) {
            vector<vector<int> > prevState = minChanges;

            // forward pass
            for (int i = 0; i < nR; i++) {
                for (int j = 0; j < nC; j++) {
                    if (i > 0) {
                        // from top
                        minChanges[i][j] = min(minChanges[i][j],
                            minChanges[i-1][j] + (grid[i-1][j] == 3 ? 0 : 1));
                    }
                    if (j > 0) {
                        // from left
                        minChanges[i][j] = min(minChanges[i][j],
                            minChanges[i][j-1] + (grid[i][j-1] == 1 ? 0 : 1));
                    }
                }
            }

            // backward pass
            for (int i = nR - 1; i >= 0; i--) {
                for (int j = nC - 1; j >= 0; j--) {
                    // from bottom
                    if (i < nR - 1) {
                        minChanges[i][j] = min(minChanges[i][j],
                            minChanges[i+1][j] + (grid[i+1][j] == 4 ? 0 : 1));
                    }
                    // from right
                    if (j < nC - 1) {
                        minChanges[i][j] = min(minChanges[i][j],
                            minChanges[i][j+1] + (grid[i][j+1] == 2 ? 0 : 1));
                    }
                }
            }

            if (prevState == minChanges) {
                break;
            }
        }
        return minChanges[nR-1][nC-1];
    }
public:
    int minCost(vector<vector<int>>& grid) {
        return solveminCost(grid);
    }
};

// Driver Code for testing
int main() {
    vector<vector<int> > grid {{1,1,1,1}, {2,2,2,2}, {1,1,1,1}, {2,2,2,2}};
    Solution *sol = new Solution();
    cout << sol->minCost(grid) << "\n";
    return 0;
}
