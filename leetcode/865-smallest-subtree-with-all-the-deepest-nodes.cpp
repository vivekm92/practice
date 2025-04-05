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

// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/
class Solution {
private:
    pair<int, TreeNode* > solveSubtreeWithAllDeepestHelper(TreeNode *root, int h) {

        if (root == nullptr) {
            return {-1, 0};
        }

        auto lp = solveSubtreeWithAllDeepestHelper(root->left, h);
        auto rp = solveSubtreeWithAllDeepestHelper(root->right, h);

        if (lp.first > rp.first) return {lp.first + 1, lp.second};
        else if (lp.first < rp.first) return {rp.first + 1, rp.second};

        return {lp.first + 1, root};
    }

    // Approach 1
    // using LCA
    // T(n) : O(n) ; S(n) : O(n)
    TreeNode* solveSubtreeWithAllDeepest(TreeNode *root) {
        return solveSubtreeWithAllDeepestHelper(root, 0).second;
    }
public:
    TreeNode* subtreeWithAllDeepest(TreeNode* root) {
        return solveSubtreeWithAllDeepest(root);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
