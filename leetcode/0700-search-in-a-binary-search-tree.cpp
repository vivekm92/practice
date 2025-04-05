#include <iostream>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// https://leetcode.com/problems/search-in-a-binary-search-tree/description/
class Solution {
private:
    // Approach 1
    // recursive approach
    // T(n) : O(n) ; S(n) : O(n)
    TreeNode* solveSearchBST(TreeNode *root, int val) {
        if (root == nullptr || root->val == val) {
            return root;
        }

        if (root->val >= val) return solveSearchBST(root->left, val);
        return solveSearchBST(root->right, val);
    }

    // Approach 2
    // iterative approach
    // T(n) : O(n) ; S(n) : O(1)
    TreeNode* solveSearchBST1(TreeNode *root, int val) {

        while (root != nullptr) {
            if (root->val == val) return root;
            else if (root->val >= val) root = root->left;
            else root = root->right;
        }

        return root;
    }
public:
    TreeNode* searchBST(TreeNode* root, int val) {
        return solveSearchBST1(root, val);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
