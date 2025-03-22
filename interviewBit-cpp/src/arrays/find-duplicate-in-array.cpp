#include <iostream>

using namespace std;

int solveFindDuplicateInArray(vector<int>& A) {
    
    int n = A.size();
    for (int i = 0; i < n; i++) {
        int idx = abs(A[i])-1;
        if (A[idx] < 0) return abs(A[i]);
        A[idx] *= -1;
    }
    
    return -1;
}

// T(n) : O(n) ; S(n) : O(n)
int findDuplicateInArray(const vector<int> &A) {
    vector<int> B = A;
    return solveFindDuplicateInArray(B);
}

// Driver Code for testing
int main() {
    
    vector<int> arr = {1,2,3,4,5,6}
    return 0;
}
