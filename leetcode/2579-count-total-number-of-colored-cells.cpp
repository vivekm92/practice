#include <iostream>

using namespace std;

class Solution {
private:
    // T(n) : O(1) ; S(n) : O(1)
    long long solveColoredCells(long long n) {
        return 2 * n * n - (2 * n - 1);
    }
public:
    long long coloredCells(int n) {
        return solveColoredCells(n);
    }
};

// Driver Code for testing
int main() {
    Solution *sol = new Solution();
    cout << sol->coloredCells(2) << "\n";
    cout << sol->coloredCells(3) << "\n";
    return 0;
}
