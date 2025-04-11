#include <iostream>
#include <queue>
#include <stack>

using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

// https://leetcode.com/problems/maximum-depth-of-binary-tree/description/
class Solution {
private:

    // Appraoch 1
    // using recursion
    // T(n) : O(n) ; S(n) : O(n)
    int solveMaxDepth(TreeNode *root) {
        if (root == nullptr) return 0;
        return max(solveMaxDepth(root->left), solveMaxDepth(root->right)) + 1;
    }

    // Approach 2
    // using bfs ( level-order traversal )
    // T(n) : O(n) ; S(n) : O(n)
    int solveMaxDepth1(TreeNode *root) {

        if (root == nullptr) return 0;

        queue<TreeNode* > q;
        q.push(root);
        int max_depth = 0;
        while (!q.empty()) {
            int sz = q.size();
            for (int i = 0; i < sz; i++) {
                root = q.front();
                q.pop();
                if (root->left != nullptr) q.push(root->left);
                if (root->right != nullptr) q.push(root->right);
            }
            max_depth++;
        }

        return max_depth;
    }

    // Approach 3
    // using stacks
    // T(n) : O(n) ; S(n) : O(n)
    int solveMaxDepth2(TreeNode* root) {

        if (root == nullptr) return 0;

        stack<pair<int, TreeNode*> > st;
        st.push({1, root});
        int max_depth = 0;
        while (!st.empty()) {
            pair<int, TreeNode* > p = st.top();
            int depth = get<0>(p);
            max_depth = max(max_depth, depth);
            st.pop();

            TreeNode *node = get<1>(p);
            if (node->left != nullptr) {
                st.push({depth + 1, node->left});
            }

            if (node->right != nullptr) {
                st.push({depth + 1, node->right});
            }
        }

        return max_depth;
    }
public:
    int maxDepth(TreeNode* root) {
        return solveMaxDepth2(root);
    }
};

// Driver Code for testing
int main() {

    return 0;
}
