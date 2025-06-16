#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/set-matrix-zeroes/description/
class Solution {
private:
    // Approach 1
    // using extra space
    // T(n) : O(m*n) ; S(n) : O(m+n)
    void solveSetZeroes(vector<vector<int> > &matrix) {
        
        int m = matrix.size(), n = matrix[0].size();
        unordered_set<int> r, c;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (matrix[i][j] == 0) {
                    r.insert(i);
                    c.insert(j);
                }
            }
        }
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (r.find(i) != r.end() || c.find(j) != c.end()) {
                    matrix[i][j] = 0;
                }
            }
        }
        
        return ;
    }
    
    // Approach 2
    // without using extra memory
    // T(n) : O(m*n) ; S(n) : O(1)
    void solveSetZeroes1(vector<vector<int> >& matrix) {
        
    }
public:
    void setZeroes(vector<vector<int>>& matrix) {
        return solveSetZeroes(matrix);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
