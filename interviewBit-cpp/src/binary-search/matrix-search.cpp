#include <iostream>

using namespace std;

// Problem : https://www.interviewbit.com/problems/matrix-search/

int matrixSearchHelper2(vector<int>&A, int B, int l, int r) {
    
}

int matrixSearch2(vector<vector<int> >& A, int B) {
    int n = A.size(), m = A[0].size();
    return matrixSearchHelper2(A, B, 0, n*m);
}
/********************************************************************************/

// Optimised Approach 2:
// T(n) : O(logn) : S(n) : O(1)
int matrixSearch2(vector<vector<int> >& A, int B) {
	
	int n = A.size(), m = A[0].size();
    int l = 0, r = m * n;
    while (l < r) {
        int mid = l + (r - l) / 2;
        if (A[mid/m][mid%m] < B) {
            l = mid + 1;
        } else {
            r = mid;
        }
    }
    return (l == m*n || A[l/m][l%m] != B) ? 0 : 1;
}
/********************************************************************************/

// Optimised Approach 1:
// T(n) : O(logn) : S(n) : O(1)
int matrixSearch1(vector<vector<int> >& A, int B) {

	int n = A.size(), m = A[0].size();
	int l = 0, r = n * m - 1;
	while (l <= r) {
		int mid = l + (r  - l) / 2;
		int val = A[mid / m][mid % m];
		if (val == B) {
			return 1;
		} else if (val < B) {
			l = mid + 1;
		} else {
			r = mid - 1;
		}
	}
	
	return 0;
}
/********************************************************************************/

// Brute-Force Approach:
// T(n) : O(n2) ; S(n) : O(1)
int matrixSearch(vector<vector<int> >& A, int B) {
	
	int n = A.size(), m = A[0].size();
	for (int i = 0; i < n; i++) { 
		for (int j = 0; j < m; j++) {
			if (A[i][j] == B) {
				return 1; 
			}
		}
	}
	
	return 0;
}
/********************************************************************************/

// Driver Code for testing
int main() {
	
	return 0;
}
