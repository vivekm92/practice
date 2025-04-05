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

// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/description/
class Solution {
private:
    pair<int, TreeNode* > solveLcaDeepestLeavesHelper(TreeNode *root, int h) {

        if (root == nullptr) {
            return {-1, root};
        }

        pair<int, TreeNode* > lp = solveLcaDeepestLeavesHelper(root->left, h);
        pair<int, TreeNode* > rp = solveLcaDeepestLeavesHelper(root->right, h);

        if (lp.first == rp.first) {
            return {lp.first + 1, root};
        } else if (lp.first > rp.first) {
            return {lp.first + 1, lp.second};
        }

        return {rp.first + 1, rp.second};
    }

    // Approach 1
    // using lca
    // T(n) : O(n) ; S(n) : O(n)
    TreeNode* solveLcaDeepestLeaves(TreeNode *root) {
        return solveLcaDeepestLeavesHelper(root, 0).second;
    }
public:
    TreeNode* lcaDeepestLeaves(TreeNode* root) {
        return solveLcaDeepestLeaves(root);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
