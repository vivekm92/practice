#include <iostream>
#include <numeric>
#include <climits>

using namespace std;

// https://leetcode.com/problems/grid-game/description/
class Solution {
private:
    // using prefix-sum
    // T(n) : O(n) ; S(n) : O(n)
    long long solvegridGame(vector<vector<int> >& grid) {
        const int n = grid[0].size();
        long long prefixSum1 = accumulate(grid[0].begin(), grid[0].end(), 0ll);
        long long minSum = LLONG_MAX, prefixSum2 = 0ll;
        for (int i = 0; i <n; i++) {
            prefixSum1 -= grid[0][i];
            minSum = min(minSum, max(prefixSum1, prefixSum2));
            prefixSum2 += grid[1][i];
        }
        return minSum;
    }
public:
    long long gridGame(vector<vector<int>>& grid) {
        return solvegridGame(grid);
    }
};

// Driver Code for testing
int main() {
    vector<vector<int> > grid = {{3,3,1}, {8,5,2}};
    Solution *sol = new Solution();
    cout << sol->gridGame(grid) << "\n";
    return 0;
}
