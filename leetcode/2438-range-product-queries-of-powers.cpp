#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/range-product-queries-of-powers/description
class Solution {
private:
	vector<int> constructPowerArray(int n) {
		
		vector<int> powerArray;
		while (n > 0) {
			int temp = 1;
			while (temp <= n) {
				temp = temp * 2;
			}
			int val = temp / 2;
			powerArray.push_back(val);
			n -= val;
		}
		
		return powerArray;
	}
	
	vector<int> solveProductQueries(int n, vector<vector<int> > &queries) {
		
		const int MOD = 1e9+7;
		vector<int> powerArray = constructPowerArray(n);
		sort(powerArray.begin(), powerArray.end());
		
		vector<int> res;
		for (auto &query : queries) { 
			int l = query[0], r = query[1];
			int prod = 1;
			for (int i = l; i <= r; i++) {
				prod = (prod * powerArray[i]) % MOD;
			}
			res.push_back(prod);
		}
		
		return res;
	}
public:
    vector<int> productQueries(int n, vector<vector<int>>& queries) {
		return solveProductQueries(n, queries);
    }
};

// Driver Code for testing
int main() {
	Solution *sol = new Solution();
	vector<vector<int> > queries = {{0,1}, {2,2}, {0,3}};
	vector<int> arr = sol->productQueries(15, queries);
	for (auto &a : arr) cout << a << " ";
	cout << "\n";
	return 0;
}
