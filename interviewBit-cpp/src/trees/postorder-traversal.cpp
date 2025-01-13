#include <iostream>
#include <vector>
#include "trees.h"

using namespace std;

void postorderTraversalRecursive(TreeNode* A, vector<int>& B) {
    
    if (A == nullptr) { 
        return ;
    }

    postorderTraversalRecursive(A->left, B);
    postorderTraversalRecursive(A->right, B);
    B.push_back(A->val);
}

vector<int> solvePostorderTraversal(TreeNode* A) {
    vector<int> traversal;
    postorderTraversalRecursive(A, traversal);
    return traversal;
}

int main() {

    return 0;
}