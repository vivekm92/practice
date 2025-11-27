#include <bits/stdc++.h>

using namespace std;

// Definition for a binary tree node.
struct TreeNode {
	int val;
	TreeNode *left;
	TreeNode *right;
	TreeNode() : val(0), left(nullptr), right(nullptr) {}
	TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
	TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// https://leetcode.com/problems/largest-bst-subtree/description/ 
class Solution {
private:
	int solveLargestBSTSubtree(TreeNode *root) {
		
	}
public:
    int largestBSTSubtree(TreeNode* root) {
        return solveLargestBSTSubtree(root);
    }
};

// Driver Code for testing
int main() {
	
	return 0;
}
