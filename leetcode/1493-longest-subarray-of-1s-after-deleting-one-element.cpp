#include <bits/stdc++.h>

using namespace std;

// https://leetcode.com/problems/longest-subarray-of-1s-after-deleting-one-element/description/
class Solution {
private:
	// Approach 1
	// using prefix and suffix array
	// T(n) : O(n) ; S(n) : O(n)
	int solveLongestSubarray(vector<int> &nums) {
		
		int n = nums.size();
		if (n <= 1) return 0;
		
		vector<int> pre(n, 0);
		pre[0] = nums[0] == 1 ? 1 : 0;
		for (int i = 1; i < n; i++) {
			if (nums[i] == 1) {
				pre[i] = pre[i-1] + 1;
			}
		}
		
		vector<int> post(n, 0);
		post[n-1] = nums[n-1] == 1 ? 1 : 0;
		for (int i = n-2; i >= 0; i--) {
			if (nums[i] == 1) {
				post[i] = post[i+1] + 1;
			}
		}
		
		int curr = 0, max_len = 0;
		for (int i = 0; i < n; i++) {
			if (i == 0) {
				curr = post[i+1];
			} else if (i == n-1) {
				curr = pre[i-1];
			} else {
				curr = pre[i-1] + post[i+1];
			}
			max_len = max(max_len, curr);
		}

		return max_len;
	}
	
	// Approach 2
	// using sliding window
	// T(n) : O(n) ; S(n): O(1)
	int solveLongestSubarray1(vector<int> &nums) {
		
		int n = nums.size();
		
	}
public:
    int longestSubarray(vector<int>& nums) {
        return solveLongestSubarray(nums);
    }
};


// Driver Code for testing.
int main() {

	return 0;
}
