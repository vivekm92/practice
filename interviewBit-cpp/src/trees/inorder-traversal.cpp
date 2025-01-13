#include <iostream>
#include <vector>
#include "trees.h"

using namespace std;

// T(n) : O(n) ; S(n) : O(n)
vector<int> inorderTraversalIterative(TreeNode* A) {

    vector<int> traversal;
    stack<TreeNode* > st;
    while (!st.empty() || A != nullptr) {
        while (A != nullptr) {
            st.push(A);
            A = A->left;
        }
        A = st.top();
        traversal.push_back(A->val);
        st.pop();
        A = A->right;
    }

    return traversal;
}

// T(n) : O(n) ; S(n) : O(n)
void inorderTraversalRecursive(TreeNode* A, vector<int>& B) {

    if (A == nullptr) {
        return ;
    }
    
    inorderTraversalRecursive(A->left, B);
    B.push_back(A->val);
    inorderTraversalRecursive(A->right, B);
}

vector<int> solveInorderTraversal(TreeNode* A) {
    vector<int> traversal;
    inorderTraversalRecursive(A, traversal);
    return traversal;
}

int main() {

    return 0;
}