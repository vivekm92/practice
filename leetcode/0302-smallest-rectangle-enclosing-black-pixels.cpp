#include <iostream>
#include <vector>

using namespace std;

// https://leetcode.com/problems/smallest-rectangle-enclosing-black-pixels/description/
class Solution {
private:
    // Approach 1
    // using naive search
    // T(n) : O(mn) ; S(n) : O(1)
    int solveMinArea(vector<vector<char> >& image, int x, int y) {
        
        int m = image.size(), n = image[0].size();
        int x1 = m, y1 = n, x2 = 0, y2 = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (image[i][j] == '1') {
                    x1 = min(x1, i);
                    x2 = max(x2, i);
                    y1 = min(y1, j);
                    y2 = max(y2, j);
                }
            }
        }
        
        return (x2-x1+1) * (y2-y1+1);
    }
    
    // Approach 2
    
public:
    int minArea(vector<vector<char>>& image, int x, int y) {
        return solveMinArea(image, x, y);
    }
};

// Driver Code for testing
int main() {
    
    return 0;
}
