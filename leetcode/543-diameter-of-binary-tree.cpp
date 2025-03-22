#include <iostream>
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

// https://leetcode.com/problems/diameter-of-binary-tree/description/
class Solution {
private:
    pair<int, int> solveDiameterOfBinaryTreeHelper(TreeNode* root) {
        
        if (root == nullptr) return {0, 0};
        
        pair<int, int> res;
        pair<int, int> lh = solveDiameterOfBinaryTreeHelper(root->left);
        pair<int, int> rh = solveDiameterOfBinaryTreeHelper(root->right);
        
        res.second = max(lh.second, rh.second) + 1;
        res.first = max(lh.first, max(rh.first, lh.second + rh.second));
        return res;
    }

    // Approach 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    int solveDiameterOfBinaryTree(TreeNode *root) {
        return solveDiameterOfBinaryTreeHelper(root).first;
    }
    
    int solveDiameterOfBinaryTreeHelper1(TreeNode *root, int &diameter) {
        
        if (root == nullptr) return 0;
        
        int lh = solveDiameterOfBinaryTreeHelper1(root->left, diameter);
        int rh = solveDiameterOfBinaryTreeHelper1(root->right, diameter);
        
        diameter = max(diameter, lh + rh);
        return max(lh, rh) + 1;
    }
    
    // Approach 2
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    int solveDiameterOfBinaryTree1(TreeNode *root) {
        
        int diameter = 0;
        solveDiameterOfBinaryTreeHelper1(root, diameter);
        return diameter;
    }
    
    // Approach 3
    // iterative approach
    // T(n) : O(n) ; S(n) : O(n)
    int solveDiameterOfBinaryTree2(TreeNode *root) {
        
        if (root == nullptr) return 0;

        TreeNode* prevNode = nullptr;
        int diameter = 0;
        unordered_map<TreeNode*, int> depths;
        depths[nullptr] = 0;
        stack<TreeNode* > st;
        while (!st.empty() || root != nullptr) {
            while (root != nullptr) {
                st.push(root);
                root = root->left;
            }
            root = st.top();
            if (root == nullptr || root == prevNode) {
                st.pop();
                diameter = max(diameter, depths[root->left] + depths[root->right]);
                depths[root] = 1 + max(depths[root->left], depths[root->right]);
                prevNode = root;
                root = nullptr;
            } else {
                root = root->right;
            }
        }
        return diameter;
    }
public:
    int diameterOfBinaryTree(TreeNode* root) {
        return solveDiameterOfBinaryTree(root);
    }
};


// Driver Code for testing
int main() {
    
    return 0;
}
