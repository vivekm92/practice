#include <iostream>
#include <vector>

using namespace std;

// https://www.interviewbit.com/problems/spiral-matrix/

// Approach 1
// T(n) : O(n) ; S(n) : O(n)
vector<vector<int> > spiralMatrix(vector<int> &A, int B, int C) {

    vector<vector<int> > spiral(B, vector<int>(C, 0));
    int rs = 0, re = B-1, cs = 0, ce = C-1, idx = 0;
    while (rs <= re && cs <= ce) {

        for (int i = cs; i <= ce; i++) {
            spiral[rs][i] = A[idx++];
        }
        rs++;
        if (cs > ce) break;

        for (int i = rs; i <= re; i++) {
            spiral[i][ce] = A[idx++];
        }
        ce--;
        if (rs > re) break;

        for (int i = ce; i >= cs; i--) {
            spiral[re][i] = A[idx++];
        }
        re--;
        if (cs > ce) break;

        for (int i = re; i >= rs; i--) {
            spiral[i][cs] = A[idx++];
        }
        cs++;
    }

    return spiral;
}

// Driver Code for testing
int main() {
    vector<int> arr = {1,2,3,4};
    vector<vector<int> > mat = spiralMatrix(arr, 2,  2);

    for (auto &vec : mat) {
        for (auto &v : vec) {
            cout << v << " ";
        }
        cout << "\n";
    }

    vector<vector<int> > mat1 = spiralMatrix(arr, 1, 4);
    for (auto &vec : mat1) {
        for (auto &v : vec) {
            cout << v << " ";
        }
        cout << "\n";
    }

    return 0;
}
