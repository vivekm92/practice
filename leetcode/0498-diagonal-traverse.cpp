#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/diagonal-traverse/description/
class Solution {
private:
	// Approach 1
	// using unordered_map
	// T(n) : O(m*n) ; S(n) : O(m*n)
	vector<int> solveFindDiagonalOrder(vector<vector<int> > &mat) {
		
		int m = mat.size(), n = mat[0].size();
		unordered_map<int, vector<int> > um;
		for (int i = 0; i < m; i++) {
			for (int j = 0; j < n; j++) {
				if (um.find(i+j) == um.end()) um[i+j] = {mat[i][j]};
				else um[i+j].push_back(mat[i][j]);
			}
		}
		
		bool flag = false;
		vector<int> res;
		for (size_t i = 0; i < um.size(); i++) {
			int k = um[i].size();
			if (!flag) {
				int x = k-1, y = 0;
				while (x >= y) {
					res.push_back(um[i][x]);
					x--;
				}
			} else {
				int x = 0, y = k-1;
				while (x <= y) {
					res.push_back(um[i][x]);
					x++;
				}
			}
			flag = flag ^ true;
		}
		
		return res;
	}
	
	// Approach 2
	// using diagonal parsing
	// T(n) : O(m*n) ; S(n) : O(m*n)
	vector<int> solveFindDiagonalOrder1(vector<vector<int> > &mat) {
		
		int m = mat.size(), n = mat[0].size();
		vector<int> res, inter;
		for (int i = 0; i < m + n - 1; i++) {
			int row, col;
			row = i<n ? 0 : i-n+1;
			col = i<n ? i : n-1;
			inter.clear();
			while (row < m && col > -1) {
				inter.push_back(mat[row][col]);
				row++;
				col--;
			}
			
			if (i%2 == 0) reverse(inter.begin(), inter.end());
			
			for (int j=0; j<inter.size(); j++) {
				res.push_back(inter[j]);
			}
		}
		
		return res;
	}
	
	// Approach 3
	// T(n) : O(m*n) ; S(n) : O(1)
	vector<int> solveFindDiagonalOrder2(vector<vector<int> > &mat) {
		
		int m = mat.size(), n = mat[0].size();
		vector<int> res;
		for (int i = 0; i < m + n - 1; i++) {
			
			
		}
		
		return res;
	}
public:
    vector<int> findDiagonalOrder(vector<vector<int>>& mat) {
        return solveFindDiagonalOrder(mat);
    }
};

// Driver Code for testing
int main() {
	vector<vector<int> > mat {{1,2,3}, {4,5,6}, {7,8,9}};
	Solution *sol = new Solution();
	vector<int> res = sol->findDiagonalOrder(mat);
	for (auto &r : res) cout << r << " ";
	cout << "\n";
	return 0;
}
