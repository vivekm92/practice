#include <iostream>
#include <vector>

using namespace std;

// https://www.interviewbit.com/problems/diagonal-flip/

// T(n) : O(n) ; S(n) : O(1)
vector<vector<int> > diagonalFlip(vector<vector<int> > &A) {

    int n = A.size();
    for (int i = 0; i < n; i++) {
        for (int j = i+1; j < n; j++) {
            swap(A[i][j], A[j][i]);
        }
    }

    return A;
}

// Driver Code for testing
int main() {

    return 0;
}
