#include <bits/stdc++.h>

using namespace std;

int solve(vector<vector<int> > &matrix) {
    
    int n = matrix.size();
    function<int(int, int)> countInversion = [&](int r, int c) {
        int count = 0;
        for (int i = r; i < n; i++) {
            for (int j = c; j < n; j++) {
                if (matrix[i][j] > matrix[r][c]) count++;
            }
        }
        return count;
    };
    
    int totalInversionsCount = 0;
    for (int i = 0; i < n; i++) {
        for (int j = 0; j < n; j++) {
            totalInversionsCount += countInversion(i, j);
        }
    }
    
    return totalInversionsCount;
}

int main() {
    
    int tc, n;
    cin >> tc;
    while (tc--) {
        cin >> n;
        vector<vector<int> > matrix(n, vector<int>(n, 0));
        for (int i = 0; i < n; i++) {
            for (int j = 0; j < n; j++) {
                cin >> matrix[i][j];
            }
        }
            
        int inversionCount = solve(matrix);
        cout << inversionCount << "\n";
    }
    return 0;
}

// inversion : matrix[i][j] > matrix[p][q] such that i <= p && j <= q
