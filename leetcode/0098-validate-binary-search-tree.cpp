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

// https://leetcode.com/problems/validate-binary-search-tree/description/
class Solution {
private:
    // Approach 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    bool solveIsValidBSTHelper(TreeNode *root, int *minVal, int *maxVal) {

        if (root == nullptr) return true;

        if ((minVal != nullptr && root->val <= *minVal) ||
            (maxVal != nullptr && root->val >= *maxVal)) return false;

        return solveIsValidBSTHelper(root->left, minVal, &(root->val)) &&
            solveIsValidBSTHelper(root->right, &(root->val), maxVal);
    }

    bool solveIsValidBST(TreeNode* root) {
        return solveIsValidBSTHelper(root, nullptr, nullptr);
    }
public:
    bool isValidBST(TreeNode* root) {
        return solveIsValidBST(root);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
