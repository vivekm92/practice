#include <bits/stdc++.h>

using namespace std;

// https://www.geeksforgeeks.org/problems/pascal-triangle0652/1
class Solution {
  private:
	int recurse(int i, int n) {
		if (i == 1 || i == n) return 1;
		return recurse(i-1, n-1) + recurse(i, n-1);
	}

	// Approach 1
	// using recursion
	// T(n) : O(2^n) ; S(n) : O(n)
	vector<int> solveNthRowOfPascalTriangle(int n) {
		
		vector<int> res;
		for (int i = 1; i <= n; i++) {
			res.push_back(recurse(i, n));
		}
		
		return res;
	}
  
    // Appraoch 2
    // using bottom-up approach
    // T(n) : O(n*n) ; S(n) : O(n)
	vector<int> solveNthRowOfPascalTriangle1(int n) {
		
		if (n == 1) return {1};
		else if (n == 2) return {1,1};
		
		vector<int> curr, prev = {1,1};
		for (int i = 3; i <= n; i++) {
			int k = prev.size();
			curr.resize(k+1, 0);
			curr[0] = 1;
			curr[k] = 1;
			for (int j = 1; j < k; j++) {
				curr[j] = prev[j-1] + prev[j];
			}
			prev = curr;
		}
		
		return curr;
	}
	
	// Approach 3
	// using level-by-level recursion
	// T(n) : O(n*n) ; S(n) : O(n)
	vector<int> solveNthRowOfPascalTriangle2(int n) {
		
		if (n == 1) return {1};
		else if (n == 2) return {1,1};
		
		vector<int> prev = solveNthRowOfPascalTriangle(n-1);
		int k = prev.size();
		vector<int> curr(k + 1, 0);
		curr[0] = 1;
		curr[k] = 1;
		for (int i = 1; i < k; i++) {
			curr[i] = prev[i-1] + prev[i];
		}
		
		return curr;
	}
	
	// Approach 4
	// using 
  public:
    vector<int> nthRowOfPascalTriangle(int n) {
        // code here
        return solveNthRowOfPascalTriangle(n);
    }
};

// Driver Code for testing.
int main() {
	
	return 0;
}
