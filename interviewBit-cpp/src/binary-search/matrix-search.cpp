#include <iostream>

using namespace std;

/*
	IB : https://www.interviewbit.com/problems/matrix-search/
	LC : 
*/

int matrixSearch2(vector<vector<int> >& A, int B) {
	
	return 1;
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
