#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/maximum-area-of-longest-diagonal-rectangle/description/
class Solution {
private:
    // Approach 1
    // using enumeration
    // T(n) : O(n) ; S(n) : O(1)
    int solveAreaOfMaxDiagonal(vector<vector<int> > &dimensions) {

        float res = -1, maxArea = -1, maxDiagonal = -1;
        for (auto &dimension : dimensions) {
            float currArea = dimension[0] * dimension[1];
            float currDiagonal = sqrt(dimension[0] * dimension[0] + dimension[1] * dimension[1]);
            if (currDiagonal > maxDiagonal || (currDiagonal == maxDiagonal && currArea > maxArea)) {
                res = currArea;
                maxArea = max(maxArea, currArea);
                maxDiagonal = max(maxDiagonal, currDiagonal);
            }
        }

        return (int )res;
    }
public:
    int areaOfMaxDiagonal(vector<vector<int>>& dimensions) {
        return solveAreaOfMaxDiagonal(dimensions);
    }
};

// Driver Code for testing.
int main() {

    return 0;
}
