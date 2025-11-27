#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/longest-harmonious-subsequence/description/
class Solution {
private:
	// Approach 1
	// brute-force approach
	// T(n) : O(2^n) ; S(n) : O(1)
	int solveFindLHS(vector<int> &nums) {
		
		int res = 0;
		int n = nums.size();
		for (int i = 0; i < (1<<n); i++) {
			int count = 0;
			long long minVal = INT_MAX, maxVal = INT_MIN;
			for (int j = 0; j < nums.size(); j++) {
				if ((i & (1 << j)) != 0) {
					minVal = min((int )minVal, nums[j]);
					maxVal = max((int )maxVal, nums[j]);
					count++;
				}
			}
			if (maxVal - minVal == 1) {
				res = max(res, count);
			}
		}
		
		return res;
	}	
public:
    int findLHS(vector<int>& nums) {
        return solveFindLHS(nums);
    }
};

// Driver Code for testing
int main() {
	
	return 0;
}
