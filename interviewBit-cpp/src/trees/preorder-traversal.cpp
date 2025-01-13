#include <iostream>
#include <vector>
#include <stack>
#include "trees.h"

using namespace std;

// T(n) : O(n) ; S(n) : O(n)
vector<int> preorderTraversalIterative(TreeNode *A) {

    vector<int> traversal;
    stack<TreeNode* > st;
    while (!st.empty() || A != nullptr) {
        while (A != nullptr) {
            traversal.push_back(A->val);
            st.push(A);
            A = A->left;
        }
        A = st.top();
        st.pop();
        A = A->right;
    }

    return traversal;
}

// T(n) : O(n) ; S(n) : O(n) 
void preorderTraversalRecursive(TreeNode* A, vector<int>& B) {
    
    if (A == nullptr) { 
        return;
    }

    B.push_back(A->val);
    preorderTraversalRecursive(A->left, B);
    preorderTraversalRecursive(A->right, B);
}

vector<int> solvePreorderTraversal(TreeNode* A) {
    vector<int> traversal;
    preorderTraversalRecursive(A, traversal);
    return traversal;
}

int main() {

    return 0;
}