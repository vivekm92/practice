#include <iostream>
#include <vector>
#include <stack>

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
 
// https://leetcode.com/problems/binary-tree-postorder-traversal/description/
class Solution {
private:
    void solvePostorderTraversalHelper(TreeNode *root, vector<int> &traversal) {
        if (root == nullptr) return ;
        solvePostorderTraversalHelper(root->left, traversal);
        solvePostorderTraversalHelper(root->right, traversal);
        traversal.push_back(root->val);
        return;
    }
    
    // Approach 1
    // recursive approach
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solvePostOrderTraversal(TreeNode *root) {
        vector<int> traversal;
        solvePostorderTraversalHelper(root, traversal);
        return traversal;
    }
    
    // Approach 2
    // iterative approach
    // T(n) : O(n) ; S(n) : O(n)
    vector<int> solvePostOrderTraversal1(TreeNode *root) {
        vector<int> traversal;
        if (root == nullptr) return traversal;
        TreeNode *prevNode = nullptr;
        stack<TreeNode* > st;
        while (!st.empty() || root != nullptr) {
            while (root != nullptr) {
                st.push(root);
                root = root->left;
            }
            root = st.top();
            if (root->right == nullptr || root->right == prevNode) {
                traversal.push_back(root->val);
                st.pop();
                prevNode = root;
                root = nullptr;
            } else {
                root = root->right;
            }
        }
        return traversal;
    }
public:
    vector<int> postorderTraversal(TreeNode* root) {
        return solvePostOrderTraversal(root);
    }
};


// Driver Code for testing
int main() {
    
    return 0;
}
